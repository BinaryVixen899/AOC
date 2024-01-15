// use libc::size_t;

// extern "C" {}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}
pub fn day1() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
