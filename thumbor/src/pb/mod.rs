use base64::{decode_config, encode_config, URL_SAFE_NO_PAD};

mod abi;
pub use abi::*;

impl ImageSpec{
    pub fn new(specs:Vec<Spec>)->Self{
        Self{specs}
    }

}
