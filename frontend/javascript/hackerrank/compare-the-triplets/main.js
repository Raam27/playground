// Compare the Triplets
// - Problem Solving (Basic)
// - Difficulty: Easy

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */


// Full Problem: https://www.hackerrank.com/challenges/compare-the-triplets/problem

function compareTriplets(a, b) {
  // Write your code here
<<<<<<< HEAD
  //beginanswer
  let arr = [0, 0]
  for (let i = 0; i < a.length; i++) {
    if (a[i] > b[i]) arr[0] = arr[0] + 1
    else if (a[i] < b[i]) arr[1] = arr[1] + 1
  }
  return arr
  //endanswer
=======
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
function main() {

  const a = [5, 6, 7] // override input
  const b = [3, 6, 10] // override input

  const result = compareTriplets(a, b);

  console.log(result)
}

main() // execute program

module.exports = compareTriplets