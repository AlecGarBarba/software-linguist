pub fn binary_search(arr: &[u32], search: u32) -> bool {
    let mut low: u32 = 0;

    let mut high: u32 = match arr.last() {
        None => 0,
        Some(t) => *t,
    };

    while low < high {
        let middle_point = low + (high - low) / 2;

        let value = match arr.get(middle_point as usize) {
            None => 0,
            Some(t) => *t,
        };

        if value == search {
            return true;
        }

        if value > search {
            high = middle_point;
        } else {
            low = middle_point + 1;
        }
    }

    return false;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn returns_false_if_empty() {
        let arr: Vec<u32> = vec![];

        assert_eq!(binary_search(&arr, 10), false);
    }

    #[test]
    fn returns_false_if_search_not_exists() {
        let arr: Vec<u32> = vec![1, 2, 3, 4, 5];

        assert_eq!(binary_search(&arr, 12389), false);
    }

    #[test]
    fn returns_true_if_beggining_array() {
        let arr: Vec<u32> = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

        assert_eq!(binary_search(&arr, *arr.first().unwrap()), true);
    }

    #[test]
    fn returns_true_if_middle_array() {
        let arr: Vec<u32> = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

        assert_eq!(binary_search(&arr, 5), true);
    }
    #[test]
    fn returns_true_if_end_array() {
        let arr: Vec<u32> = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

        assert_eq!(binary_search(&arr, *arr.last().unwrap()), true);
    }
}
