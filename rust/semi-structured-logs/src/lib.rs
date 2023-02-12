// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

/// various log levels
#[derive(Clone, PartialEq, Eq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
    Debug,
}
/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    match level {
        LogLevel::Info => info(message),
        LogLevel::Warning => warn(message),
        LogLevel::Error => error(message),
        LogLevel::Debug => debug(message),
    }
}

fn debug(message: &str) -> String {
    let mut debug_label: String = "[DEBUG]: ".to_owned();
    debug_label.push_str(message);
    return debug_label;
}
pub fn info(message: &str) -> String {
    let mut info_label: String = "[INFO]: ".to_owned();
    info_label.push_str(message);
    return info_label;
}
pub fn warn(message: &str) -> String {
    let mut warn_label: String = "[WARNING]: ".to_owned();
    warn_label.push_str(message);
    return warn_label;
}
pub fn error(message: &str) -> String {
    let mut error_label: String = "[ERROR]: ".to_owned();
    error_label.push_str(message);
    return error_label;
}
