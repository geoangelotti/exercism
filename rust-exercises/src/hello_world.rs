// &'static is a "lifetime specifier", something you'll learn more about later
fn hello() -> &'static str {
    "Hello, World!"
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hello_world() {
        assert_eq!("Hello, World!", hello());
    }
}
