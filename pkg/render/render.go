/*package render
import(
	"net/http"
	"html/template"
	"log"
)

func RenderTemplateTest(w http.ResponseWriter,tmpl string){
	parsedTemplate,_:=template.ParseFiles("./templates/"+tmpl,"./templates/base.layout.html")
	err:=parsedTemplate.Execute(w,nil)
	if err!=nil{
		log.Println("Error parsing the template")
		return
	}
}

var tc =make(map[string]*template.Template )

func RenderTemplate(w http.ResponseWriter,t string){

	var tmpl *template.Template
	var err error
	_,inMap:=tc[t]
	if !inMap{
		//need to create template
		log.Println("creating template and adding to cache")
		err=createTemplateCache(t)
		if err!=nil{
			log.Println(err)
		}

	} else {
		//template in cache
		log.Println("using cached template")
	}
	tmpl=tc[t]
	err=tmpl.Execute(w,nil)
	if err!=nil{
		log.Println(err)
	}
}
func createTemplateCache(t string)(error){
	templates:=[]string{
		"./templates/"+t,
		"./templates/base.layout.html",
	}
	tmpl,err:=template.ParseFiles(templates...)
	if err!=nil{
		return err
	}
	tc[t]=tmpl
	return nil
}*/

// Second template cache method
package render

import (
	"bytes"
	"github.com/pauljc-probeplus/bookings/pkg/config"
	//"github.com/pauljc-probeplus/lesson-5/pkg/handlers"
	"github.com/pauljc-probeplus/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions=template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app=a
}


func AddDefaultData(td *models.TemplateData)*models.TemplateData{
	return td
}


func RenderTemplate(w http.ResponseWriter,tmpl string,td *models.TemplateData){
	//get the template cache from app config
	var template_cache map[string]*template.Template
	if app.UseCache{
		template_cache=app.TemplateCache
	}else{
		template_cache,_=CreateTemplateCache()
	}
	//get requested template from cache
	template,ok:=template_cache[tmpl]
	if !ok{
		log.Fatal("could not get template from template cache")
	}

	buf:=new(bytes.Buffer)


	td=AddDefaultData(td)
	_=template.Execute(buf,td)
	
	//render the template
	
	_,err:=buf.WriteTo(w)
	if err!=nil{
		log.Println(err)
	}
}

func CreateTemplateCache()(map[string]*template.Template,error){
	myCache:=map[string]*template.Template{}

	//get all of the files named .page.html from templates folder
	pages,err:=filepath.Glob("./templates/*.page.html")
	if err!=nil{
		return myCache,err
	}

	//range through all files ending with *.page.html
	for _,page:=range(pages){
		name:=filepath.Base(page)
		template_set,err:=template.New(name).ParseFiles(page)
		if err!=nil{
			return myCache,err
		}

		matches,err:=filepath.Glob("./templates/*.layout.html")
		if err!=nil{
			return myCache,err
		}

		if len(matches)>0{
			template_set,err =template_set.ParseGlob("./templates/*.layout.html")
			if err!=nil{
				return myCache,err
			}
		}
		myCache[name]=template_set
	}
	return myCache,err
}
