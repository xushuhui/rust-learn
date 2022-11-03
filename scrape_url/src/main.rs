
use std::{fs };


fn main(){
   let mut args = Vec::new();
   for (k, arg) in std::env::args().enumerate(){
    if k > 0 {
       args.push(arg);
    }
  
   }
   println!("print {:?}",args);
    spider(&args[0], &args[1]) 
}
fn spider_with_error()->Result<(),Box<dyn std::error::Error>>{
   let url = "https://www.rust-lang.org/";
   let output = "rust.md"; 
   let body = reqwest::blocking::get(url)?.text()?;
   let md = html2md::parse_html(&body);
   fs::write(output, md.as_bytes()).unwrap();
   Ok(())
}
fn spider(url: &String,output: &String){

   let body = reqwest::blocking::get(url).unwrap().text().unwrap();
   let md = html2md::parse_html(&body);
   fs::write(output, md.as_bytes()).unwrap();
}