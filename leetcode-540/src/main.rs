pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
    let (mut begin, mut end) = (0, nums.len() - 1);
    loop {
        if begin == end {
            return nums[begin];
        }
        let mid = begin + ((end - begin) >> 1);
        // println!("begin: {} end: {} mid: {}", begin, end, mid);
        let n = mid - begin;
        if n % 2 == 0 {
            // 前面有偶数个数字， 与前一个数字相等，单个数字在前面
            if nums[mid] == nums[mid - 1] {
                end = mid - 2;
            } else {
                begin = mid;
            }
        } else {
            // 前面有奇数个数字， 与前一个数字相等，单个数字在后面
            if nums[mid] == nums[mid - 1] {
                begin = mid + 1;
            } else {
                end = mid - 1;
            }
        }
    }
}

fn main() {
    println!("Hello, world!");
    println!("{} {:?}", 0, single_non_duplicate(vec![0]));
    println!("{} {:?}", 1, single_non_duplicate(vec![0, 0, 1]));
    println!("{} {:?}", 0, single_non_duplicate(vec![0, 1, 1]));
    println!("{} {:?}", 2, single_non_duplicate(vec![0, 0, 1, 1, 2]));
    println!("{} {:?}", 1, single_non_duplicate(vec![0, 0, 1, 2, 2]));
    println!("{} {:?}", 0, single_non_duplicate(vec![0, 1, 1, 2, 2]));
    println!(
        "{} {:?}",
        3,
        single_non_duplicate(vec![0, 0, 1, 1, 2, 2, 3])
    );
    println!(
        "{} {:?}",
        2,
        single_non_duplicate(vec![0, 0, 1, 1, 2, 3, 3])
    );
    println!(
        "{} {:?}",
        1,
        single_non_duplicate(vec![0, 0, 1, 2, 2, 3, 3])
    );
    println!(
        "{} {:?}",
        0,
        single_non_duplicate(vec![0, 1, 1, 2, 2, 3, 3])
    );
}
