use clap::Parser ;
use colored::Colorize;
use mime::Mime;
use reqwest::{header,Url,Client, Response};
use anyhow::{anyhow,Result, Ok};
use std::{str::FromStr, collections::HashMap};
#[derive(Parser,Debug)]
#[clap(version="1.0",author="author")]
struct Opts{
    #[clap(subcommand)]
    subcmd: SubCommand,
}
#[derive(Parser,Debug)]
enum SubCommand {
    Get(Get),
    Post(Post),
}
#[derive(Parser,Debug)]
struct Get{
    #[clap(parse(try_from_str=parse_url))]
    url:String,
}

fn print_headers(resp:&Response){
    for (name,value) in resp.headers(){
        println!("{}: {:?}",name.to_string().green(),value);
 
    }
    print!("\n")
}
fn print_status(resp:&Response){
    let status = format!("{:?}{}",resp.version(),resp.status()).blue();
    println!("{}\n",status)
}
fn print_body(m:Option<Mime>,body:&String){
    match m {
        Some(v) if v == mime::APPLICATION_JSON=>{
            println!("{}",jsonxf::pretty_print(body).unwrap().cyan())
        }
        _=>println!("{}",body)
    }
}
async fn print_resp(resp:Response)->Result<()>{
    print_status(&resp);
    print_headers(&resp);
    let mime = get_content_type(&resp);
    let body = resp.text().await.unwrap();
    print_body(mime, &body);
    Ok(())
}
fn get_content_type(resp:&Response)->Option<Mime>{
    resp.headers().get(header::CONTENT_TYPE).map(|v|v.to_str().unwrap().parse().unwrap())
}
#[derive(Parser,Debug)]
struct Post{
    #[clap(parse(try_from_str=parse_url))]
   url:String,
   #[clap(parse(try_from_str=parse_kv_pair))]
   body:Vec<KvPair>, 
}
#[derive(Debug)]
struct KvPair{
    k:String,
    v:String,
}
impl FromStr for KvPair {
    type Err = anyhow::Error;
    fn from_str(s: &str)->Result<Self,Self::Err>{
        let mut split = s.split("=");
        let err = || anyhow!(format!("Failed to parse {}",s));
        Ok(Self { k: (split.next().ok_or_else(err)?).to_string(), v: (split.next().ok_or_else(err)?).to_string() })
    }
}
fn parse_kv_pair(s:&str)->Result<KvPair>{
    Ok(s.parse()?)
}
fn parse_url(s: &str)->Result<String>{
    let _url:Url = s.parse()?;
    Ok(s.into())
}
async fn get(client:Client,args:&Get)->Result<()>{
    let resp = client.get(&args.url).send().await?;
  
    Ok(print_resp(resp).await?)
}
async fn post(client:Client,args:&Post)->Result<()>{
    let mut body = HashMap::new();
    for pair in args.body.iter(){
        body.insert(&pair.k, &pair.v);
    }
    let resp = client.post(&args.url).json(&body).send().await?;
    Ok(print_resp(resp).await?)
    
}
#[tokio::main]
async fn main()->Result<()> {
    let opts: Opts = Opts::parse();
    let mut headers = header::HeaderMap::new();
    headers.insert("X_PWOERED_BY", "Rust".parse()?);

    let client = Client::builder().default_headers(headers).build()?;

    let result = match opts.subcmd {
        SubCommand::Get(ref args)=> get(client,args).await?,
        SubCommand::Post(ref args)=>post(client, args).await?,
    };

    Ok(result)
}

