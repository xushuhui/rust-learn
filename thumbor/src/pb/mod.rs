use base64::{decode_config, encode_config, URL_SAFE_NO_PAD};
use photon_rs::transform::SamplingFilter;
mod abi;
pub use abi::*;
use prost::Message;

impl ImageSpec{
    pub fn new(specs:Vec<Spec>)->Self{
        Self{specs}
    }
}
impl From<&ImageSpec> for String{
    fn from(image_spec:&ImageSpec)->Self{
        let data = image_spec.encode_to_vec();
        encode_config(data,URL_SAFE_NO_PAD)
    }
}
impl TryFrom<&str> for ImageSpec{
    type Error = anyhow::Error;
    fn try_from(value:&str)->Result<Self,Self::Error>{
        let data = decode_config(value,URL_SAFE_NO_PAD)?;
        Ok(ImageSpec::decode(&data[..])?)
    }
}
impl filter::Filter{
    pub fn to_str(&self)->Option<&'static str> {
        match self {
            filter::Filter::Unspecified=>None,
            filter::Filter::Oceanic=>Some("oceanic"),
            filter::Filter::Islands=>Some("islands"),
            filter::Filter::Marine=>Some("marine")
        }
    }
}
impl From<resize::SampleFilter> for SamplingFilter {
    fn from(v: resize::SampleFilter) -> Self {
        match v {
            resize::SampleFilter::Undefined=>SamplingFilter::Nearest,
            resize::SampleFilter::Nearest =>SamplingFilter::Nearest,
            resize::SampleFilter::Lanczos3 =>SamplingFilter::Lanczos3,
            resize::SampleFilter::CatmullRom=>SamplingFilter::CatmullRom,
            resize::SampleFilter::Gaussian=>SamplingFilter::Gaussian,
            resize::SampleFilter::Triangle=>SamplingFilter::Triangle,
        }
    }
}
