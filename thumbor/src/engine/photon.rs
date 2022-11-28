use super::{Engine,SpecTransform};
use crate::pb::*;
use anyhow::Result;
use bytes::Bytes;
use image::{DynamicImage,ImageBuffer,ImageOutputFormat};
use lazy_static::lazy_static;
use photon_rs::{
    effects,filters,multiple,native::open_image_from_bytes,transform,PhotonImage,
};
use std::convert::TryFrom;
lazy_Static!{
    static ref WATEMARK: PhotonImage = {
        let data = include_bytes("../../rust-logo.png");
        let watermark = open_image_from_bytes(data).unwrap();
        transform::resize(&watermark,64,64,transform::SamplingFilter::Nearest)
    }
}
pub struct Photon(PhotonImage);
impl TryFrom<Bytes> for Photon{
    type Error = anyhow::Error;
    fn try_from(data:Bytes)->Result<Self,Self::Error>{
        Ok(Self(open_image_from_bytes(&data)?))
    }
}
impl Engine for Photon{
    fn apply(&mut self,specs: &[Spec]){
        for spec in specs.iter(){
            match spec.data{
                Some(spec::Data::Crop(ref v))=> self.transform(v),
            }
        }
    }
    fn generate(self,format: ImageOutputFormat)->Vec<u8>{
        
    }
}