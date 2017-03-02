/*day 1 of week 8 assignment
*work started 11:36, 03.02.2017
*/

package main

import(
"net/http"
"html/template"
)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob("templates/*.gohtml")
}

func main(){
http.HandleFunc("/", index1)
http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))
http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
http.ExecuteTemplate(w, "index1.gohtml")
}