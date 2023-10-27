mod binary_search;

fn main() {
    let vec: Vec<u32> = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    let exists = binary_search::binary_search(&vec, 10);

    println!("{exists}");
}
