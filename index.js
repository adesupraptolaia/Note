/*
Bikin program yang melakukan looping dari 1 sampai 100
Dengan ketentuan
1. Jika bilangann tsb habis dibagi 3 maka outputnya Fizz
2. Jika bilangann tsb habis dibagi 5 maka outputnya Buzz
3. Jika bilangann tsb habis dibagi 3 dan 5 sekaligus maka outputnya FizzBuzz
*/

// jika genap tampilin genap
// jika ganjil tampilin ganjil

for (let i = 1; i <= 100; i++) {
    // jika i habis dibagi 2 maka genap
    if (i % 3 === 0) {
        console.log(i, "Fizz")
    } else if (i % 5 === 0) {
        console.log(i, "Buzz")
    }
}

/*
contoh ouput
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
dst
*/