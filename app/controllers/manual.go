package controllers

import (
	"github.com/robfig/revel"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"html/template"
	"strings"
)
type Manual struct {
	*revel.Controller
}


func (c Manual) Index() revel.Result {

	sub_dir := c.Params.Get("dir")
	filename := c.Params.Get("filename")
	root_dir := "src/help.rzj.me/manual/"
	
	cur_dir := root_dir + sub_dir
	
	var filepath string
	if filename=="" {
		_,err := ioutil.ReadFile(cur_dir+"index.md")
		if err==nil{
			filepath = "index.md"
		}else{
			filepath = ""
		}
	}else{
		filepath =filename
	}
	
	if filepath =="" {
         files, err := ioutil.ReadDir(cur_dir)
         if err!=nil {
         	return c.RenderError(err)
         }
 
         file_arr := make(map[int]string) 
         dir_arr := make(map[int]string) 
         for i, f := range files { 
            if strings.Index(f.Name(),".")==0 {
                    continue                    
            }
             if !f.IsDir() {
             	file_arr[i] = f.Name()
             } else {
             	dir_arr[i] =f.Name()
             }
         }
		 return c.Render(dir_arr,file_arr)
	}
	
	cur_path :=cur_dir + filepath
	content,err := ioutil.ReadFile(cur_path)
	if err!=nil {
		return c.RenderTemplate("errors/404.html")
	}
	output := template.HTML(string(blackfriday.MarkdownBasic(content)))
    return c.Render(output)
}
