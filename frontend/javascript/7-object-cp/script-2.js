// Buat kalkulator objek dengan 5 methods:
// - read() meminta (prompt) dua nilai dan menyimpannya sebagai properti objek.
// - sum() mengembalikan jumlah nilai yang disimpan.
// - substract() mengurangi jumlah nilai yang disimpan.
// - multiply() mengalikan nilai yang disimpan.
// - division() membagi nilai yang disimpan.

let calculator = {
  // Tulis kode di sini
<<<<<<< HEAD
  // beginanswer
  sum() {
    return this.a + this.b;
  },

  substract() {
    return this.a - this.b;
  },

  multiply() {
    return this.a * this.b;
  },

  division() {
    return this.a / this.b;
  },

  read() {
    this.a = +prompt('a =', 0);
    this.b = +prompt('b =', 0);
  }
  // endanswer
=======
  // TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
};
  
calculator.read();
console.log(calculator.sum());
console.log(calculator.substract());
console.log(calculator.multiply());
console.log(calculator.division());