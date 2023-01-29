defmodule ReverseList do

  def reverse(list), do: reverse(list, [])

  def reverse([head | tail], acc), do: reverse(tail, [head | acc])

  def reverse([], acc), do: acc
end

ReverseList.reverse([1,2,3,4,5,6,7,8,9])
|> IO.inspect()
