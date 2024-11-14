function binarySearch(numberList, item) {
    let lowNumber = 0;
    let highNumber = numberList.length - 1;

    while (lowNumber <= highNumber) {
        let mid = Math.floor((lowNumber + highNumber) / 2);

        if (numberList[mid] === item) {
            return mid;
        } else if (numberList[mid] < item) {
            lowNumber = mid + 1;
        } else {
            highNumber = mid - 1;
        }
    }
}

let myList = [1, 3, 5, 7, 9];

console.log(binarySearch(myList, 3));