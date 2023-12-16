package math

func MinIntSlice(s []int) int {
  m := 0

  if len(s) == 0 {
    return 0
  }

  for i, e := range s {
    if i==0 || e < m {
      m = e
    }
  }
  return m
}

func Sum(l []int) int {
  s := 0

  for _, e := range l {
    s += e
  }
  return s
}
