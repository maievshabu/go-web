package main
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)
type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title  + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil{
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hi hell world %s and %s", r.URL.Path[1:], r.Method)
}

func viewHanler(w http.ResponseWriter, r *http.Request){

	title := r.URL.Path[len("/view/"):]
	p,_ := loadPage(title)

	fmt.Fprintf(w, "hi hell world %s and %s", p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func main(){
	http.HandleFunc("/hello/", handler)
	http.HandleFunc("/view/", viewHanler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":7896", nil))

	//p1 := &Page{Title:"testPage", Body:[]byte("this he llo")}
	//p1.save()
	//
	//p2,_ := loadPage("testPage")
	//fmt.Println(string(p2.Body));
}
