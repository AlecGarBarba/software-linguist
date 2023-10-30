pub fn ascending(a: i32, b: i32) -> bool {
    a > b
}

pub fn descending(a: i32, b: i32) -> bool {
    a < b
}

pub fn bubblesort(arr: &mut [i32], should_order: fn(a: i32, b: i32) -> bool) -> () {
    if arr.len() == 0 {
        return ();
    }
    for i in 1..=arr.len() - 1 {
        for j in 0..=arr.len() - i - 1 {
            let curr = arr[j];
            let next = arr[j + 1];
            if should_order(curr, next) {
                arr[j + 1] = curr;
                arr[j] = next;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_orders_ascending() {
        let mut arr: Vec<i32> = vec![1, 3, 7, 2, 4, 6, 5, 8];
        bubblesort(&mut arr, ascending);
        assert_eq!(arr, vec![1, 2, 3, 4, 5, 6, 7, 8]);
    }

    #[test]
    fn it_orders_descending() {
        let mut arr: Vec<i32> = vec![1, 3, 7, 2, 4, 6, 5, 8];
        bubblesort(&mut arr, descending);
        assert_eq!(arr, vec![8, 7, 6, 5, 4, 3, 2, 1]);
    }

    #[test]
    fn it_does_not_panic_on_empty() {
        let mut arr: Vec<i32> = Vec::new();
        bubblesort(&mut arr, ascending);
        assert_eq!(arr, Vec::new());
    }
}
