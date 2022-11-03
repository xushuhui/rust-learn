#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ImageSpec {
    #[prost(message, repeated, tag="1")]
    pub specs: ::prost::alloc::vec::Vec<Spec>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Spec {
    #[prost(oneof="spec::Data", tags="1, 2, 3, 4, 5, 6, 7")]
    pub data: ::core::option::Option<spec::Data>,
}
/// Nested message and enum types in `Spec`.
pub mod spec {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Data {
        #[prost(message, tag="1")]
        Resize(super::Resize),
        #[prost(message, tag="2")]
        Crop(super::Crop),
        #[prost(message, tag="3")]
        Flipv(super::Flipv),
        #[prost(message, tag="4")]
        Fliph(super::Fliph),
        #[prost(message, tag="5")]
        Contrast(super::Contrast),
        #[prost(message, tag="6")]
        Filter(super::Filter),
        #[prost(message, tag="7")]
        Watermark(super::Watermark),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Resize {
    #[prost(uint32, tag="1")]
    pub width: u32,
    #[prost(uint32, tag="2")]
    pub hight: u32,
    #[prost(enumeration="resize::ResizeType", tag="3")]
    pub r_type: i32,
    #[prost(enumeration="resize::SampleFilter", tag="4")]
    pub filter: i32,
}
/// Nested message and enum types in `Resize`.
pub mod resize {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum ResizeType {
        Normal = 0,
        SeamCarve = 1,
    }
    impl ResizeType {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                ResizeType::Normal => "NORMAL",
                ResizeType::SeamCarve => "SEAM_CARVE",
            }
        }
    }
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum SampleFilter {
        Undefined = 0,
        Nearest = 1,
        Triangle = 2,
        CatmullRom = 3,
        Gaussian = 4,
        Lanczos3 = 5,
    }
    impl SampleFilter {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                SampleFilter::Undefined => "UNDEFINED",
                SampleFilter::Nearest => "NEAREST",
                SampleFilter::Triangle => "TRIANGLE",
                SampleFilter::CatmullRom => "CATMULL_ROM",
                SampleFilter::Gaussian => "GAUSSIAN",
                SampleFilter::Lanczos3 => "LANCZOS3",
            }
        }
    }
}
/// 处理图片截取
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Crop {
    #[prost(uint32, tag="1")]
    pub x1: u32,
    #[prost(uint32, tag="2")]
    pub y1: u32,
    #[prost(uint32, tag="3")]
    pub x2: u32,
    #[prost(uint32, tag="4")]
    pub y2: u32,
}
/// 处理水平翻转
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Fliph {
}
/// 处理垂直翻转
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Flipv {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Contrast {
    #[prost(float, tag="1")]
    pub contrast: f32,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Filter {
    #[prost(enumeration="filter::Filter", tag="1")]
    pub filter: i32,
}
/// Nested message and enum types in `Filter`.
pub mod filter {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Filter {
        Unspecified = 0,
        Oceanic = 1,
        Islands = 2,
        /// more: <https://docs.rs/photon-rs/0.3.1/photon_rs/filters/fn.filter.html> 
        Marine = 3,
    }
    impl Filter {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Filter::Unspecified => "UNSPECIFIED",
                Filter::Oceanic => "OCEANIC",
                Filter::Islands => "ISLANDS",
                Filter::Marine => "MARINE",
            }
        }
    }
}
/// 处理水印
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Watermark {
    #[prost(uint32, tag="1")]
    pub x: u32,
    #[prost(uint32, tag="2")]
    pub y: u32,
}
