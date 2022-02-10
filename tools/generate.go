package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/linden/chroma"
	"github.com/linden/chroma/formatters/html"
	"github.com/linden/chroma/lexers"
	"github.com/linden/chroma/styles"

	"github.com/russross/blackfriday/v2"
)

// Segment is a segment of an example
type Segment struct {
	Docs         string
	DocsRendered string
	Code         string
	CodeRendered string
	CodeForJs    string
	CodeEmpty    bool
	CodeLeading  bool
	CodeRun      bool
}

// Example is info extracted from an example file
type Example struct {
	ID          string
	Name        string
	GoCode      string
	GoCodeHash  string
	URLHash     string
	Segments    [][]*Segment
	PrevExample *Example
	NextExample *Example
}

// siteDir is the target directory into which the HTML gets generated. Its
// default is set here but can be changed by an argument passed into the
// program.
var siteDir = "./public"

func verbose() bool {
	return len(os.Getenv("VERBOSE")) > 0
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ensureDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	check(err)
}

func copyFile(src, dst string) {
	dat, err := os.ReadFile(src)
	check(err)

	err = os.WriteFile(dst, dat, 0644)
	check(err)
}

func pipe(bin string, arg []string, src string) []byte {
	cmd := exec.Command(bin, arg...)

	in, err := cmd.StdinPipe()
	check(err)

	out, err := cmd.StdoutPipe()
	check(err)

	err = cmd.Start()
	check(err)

	_, err = in.Write([]byte(src))
	check(err)

	err = in.Close()
	check(err)

	bytes, err := io.ReadAll(out)
	check(err)

	err = cmd.Wait()
	check(err)

	return bytes
}

func sha1Sum(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func mustReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	check(err)

	return string(bytes)
}

func markdown(src string) string {
	return string(blackfriday.Run([]byte(src)))
}

func readLines(path string) []string {
	src := mustReadFile(path)
	return strings.Split(src, "\n")
}

func mustGlob(glob string) []string {
	paths, err := filepath.Glob(glob)
	check(err)

	return paths
}

func debug(msg string) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Fprintln(os.Stderr, msg)
	}
}

var docsPat = regexp.MustCompile("^\\s*(\\/\\/|#|;;)\\s")
var dashPat = regexp.MustCompile("\\-+")

func parseHashFile(sourcePath string) (string, string) {
	lines := readLines(sourcePath)
	return lines[0], lines[1]
}

func parseSegments(sourcePath string) ([]*Segment, string) {
	var (
		lines  []string
		source []string
	)

	// Convert tabs to spaces for uniform rendering.
	for _, line := range readLines(sourcePath) {
		lines = append(lines, strings.Replace(line, "\t", "    ", -1))
		source = append(source, line)
	}

	filecontent := strings.Join(source, "\n")
	segments := []*Segment{}
	lastSeen := ""

	for _, line := range lines {
		if line == "" {
			lastSeen = ""
			continue
		}

		matchDocs := docsPat.MatchString(line)
		matchCode := !matchDocs

		newDocs := (lastSeen == "") || ((lastSeen != "docs") && (segments[len(segments)-1].Docs != ""))
		newCode := (lastSeen == "") || ((lastSeen != "code") && (segments[len(segments)-1].Code != ""))

		if newDocs || newCode {
			debug("NEWSEG")
		}

		if matchDocs {
			trimmed := docsPat.ReplaceAllString(line, "")

			if newDocs {
				newSegment := Segment{Docs: trimmed, Code: ""}
				segments = append(segments, &newSegment)
			} else {
				segments[len(segments)-1].Docs = segments[len(segments)-1].Docs + "\n" + trimmed
			}

			debug("DOCS: " + line)

			lastSeen = "docs"
		} else if matchCode {
			if newCode {
				newSegment := Segment{Docs: "", Code: line}
				segments = append(segments, &newSegment)
			} else {
				segments[len(segments)-1].Code = segments[len(segments)-1].Code + "\n" + line
			}

			debug("CODE: " + line)

			lastSeen = "code"
		}
	}

	for index, segment := range segments {
		segment.CodeEmpty = (segment.Code == "")
		segment.CodeLeading = (index < (len(segments) - 1))
		segment.CodeRun = strings.Contains(segment.Code, "package main")
	}

	return segments, filecontent
}

func chromaFormat(code string, filePath string) string {
	lexer := lexers.Get(filePath)

	if lexer == nil {
		lexer = lexers.Fallback
	}

	if strings.HasSuffix(filePath, ".sh") {
		lexer = SimpleShellOutputLexer
	}

	lexer = chroma.Coalesce(lexer)

	style := styles.Get("swapoff")

	if style == nil {
		style = styles.Fallback
	}

	formatter := html.New(html.WithClasses(true))
	iterator, err := lexer.Tokenise(nil, string(code))

	check(err)

	buf := new(bytes.Buffer)

	err = formatter.Format(buf, style, iterator)
	check(err)

	return buf.String()
}

func parseAndRenderSegments(sourcePath string) ([]*Segment, string) {
	segments, filecontent := parseSegments(sourcePath)

	for _, segment := range segments {
		if segment.Docs != "" {
			segment.DocsRendered = markdown(segment.Docs)
		}

		if segment.Code != "" {
			segment.CodeRendered = chromaFormat(segment.Code, sourcePath)
		}
	}

	return segments, filecontent
}

func parseExamples() []*Example {
	var exampleNames []string

	for _, line := range readLines("examples.txt") {
		if line != "" && !strings.HasPrefix(line, "#") {
			exampleNames = append(exampleNames, line)
		}
	}

	examples := make([]*Example, 0)

	for index, exampleName := range exampleNames {
		if verbose() {
			fmt.Printf("Processing %s [%d/%d]\n", exampleName, index+1, len(exampleNames))
		}

		var editorializedName string

		if exampleName == "SIP-10" {
			editorializedName = "SIP-10 (Tokens)"
		} else {
			editorializedName = exampleName
		}

		example := Example{Name: editorializedName}
		exampleID := strings.ToLower(exampleName)
		exampleID = strings.Replace(exampleID, " ", "-", -1)
		exampleID = strings.Replace(exampleID, "/", "-", -1)
		exampleID = strings.Replace(exampleID, "'", "", -1)
		exampleID = dashPat.ReplaceAllString(exampleID, "-")
		example.ID = exampleID
		example.Segments = make([][]*Segment, 0)

		sourcePaths := mustGlob("examples/" + exampleID + "/*")

		for _, sourcePath := range sourcePaths {
			sourceSegments, filecontents := parseAndRenderSegments(sourcePath)

			if filecontents != "" {
				example.GoCode = filecontents
			}

			example.Segments = append(example.Segments, sourceSegments)
		}

		examples = append(examples, &example)
	}

	for index, example := range examples {
		if index > 0 {
			example.PrevExample = examples[index-1]
		}
		if index < (len(examples) - 1) {
			example.NextExample = examples[index+1]
		}
	}

	return examples
}

func renderIndex(examples []*Example) {
	if verbose() {
		fmt.Println("Rendering index")
	}

	indexTmpl := template.New("index")

	_, err := indexTmpl.Parse(mustReadFile("templates/footer.tmpl"))
	check(err)

	_, err = indexTmpl.Parse(mustReadFile("templates/index.tmpl"))
	check(err)

	indexF, err := os.Create(siteDir + "/index.html")
	check(err)

	err = indexTmpl.Execute(indexF, examples)
	check(err)
}

func renderExamples(examples []*Example) {
	if verbose() {
		fmt.Println("Rendering examples")
	}

	exampleTmpl := template.New("example")

	_, err := exampleTmpl.Parse(mustReadFile("templates/footer.tmpl"))
	check(err)

	_, err = exampleTmpl.Parse(mustReadFile("templates/example.tmpl"))
	check(err)

	for _, example := range examples {
		exampleF, err := os.Create(siteDir + "/" + example.ID + ".html")
		check(err)

		exampleTmpl.Execute(exampleF, example)
	}
}

func main() {
	if len(os.Args) > 1 {
		siteDir = os.Args[1]
	}

	ensureDir(siteDir)

	copyFile("templates/site.css", siteDir+"/site.css")
	copyFile("templates/site.js", siteDir+"/site.js")
	copyFile("templates/favicon.ico", siteDir+"/favicon.ico")
	copyFile("templates/404.html", siteDir+"/404.html")
	copyFile("templates/clipboard.png", siteDir+"/clipboard.png")

	examples := parseExamples()

	renderIndex(examples)
	renderExamples(examples)
}

var SimpleShellOutputLexer = chroma.MustNewLexer(
	&chroma.Config{
		Name:      "Shell Output",
		Aliases:   []string{"console"},
		Filenames: []string{"*.sh"},
		MimeTypes: []string{},
	},
	func() chroma.Rules {
		return chroma.Rules{
			"root": {
				// $ or > triggers the start of prompt formatting
				{`^\$`, chroma.GenericPrompt, chroma.Push("prompt")},
				{`^>`, chroma.GenericPrompt, chroma.Push("prompt")},

				// empty lines are just text
				{`^$\n`, chroma.Text, nil},

				// otherwise its all output
				{`[^\n]+$\n?`, chroma.GenericOutput, nil},
			},
			"prompt": {
				// when we find newline, do output formatting rules
				{`\n`, chroma.Text, chroma.Push("output")},
				// otherwise its all text
				{`[^\n]+$`, chroma.Text, nil},
			},
			"output": {
				// sometimes there isn't output so we go right back to prompt
				{`^\$`, chroma.GenericPrompt, chroma.Pop(1)},
				{`^>`, chroma.GenericPrompt, chroma.Pop(1)},
				// otherwise its all output
				{`[^\n]+$\n?`, chroma.GenericOutput, nil},
			},
		}
	},
)
