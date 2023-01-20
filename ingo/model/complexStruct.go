package model

type ComplexStruct struct {
  Number int `firestore:"number"`
  Fib int `firestore:"fib"`
  Sum int `firestore:"sum"`
  Random string `firestore:"random"`
  Parity bool `firestore:"parity"`
}
