# Постановка задачи

# Нам дают упорядоченный массив и определенное значение. Затем просят создать функцию, которая возвращает true или false, в зависимости от того, может ли сумма любых двух чисел из массива быть равной заданному значению.

# Другими словами, есть ли в массиве два целых числа x и y, которые при сложении равны указанному значению?

# Пример А

# Если нам дали массив [1, 2, 4, 9] и значение 8, функция вернет false, потому что никакие два числа из массива не могут дать 8 в сумме.

# Пример Б

# Но если это массив [1, 2, 4, 4] и значение 8, функция должна вернуть true, потому что 4 + 4 = 8.

defmodule SumTwoNumbersInList do
  def is_two_numbers_exist?([], _k), do: false
  def is_two_numbers_exist?(list, _k) when length(list) in [0, 1], do: false

  def is_two_numbers_exist?(list, k) do
    linear_search = linear_search(list, k)

    binary_search =
      Enum.reduce_while(list, list, fn h, [h | t] ->
        # IO.inspect({h, t, k-h})

        cond do
          length(t) == 0 ->
            {:halt, false}

          binary_search(t, k - h) ->
            {:halt, true}

          true ->
            {:cont, t}
        end
      end)

    {binary_search, linear_search}
  end

  # HELPERS

  # binary search
  defp binary_search(list, _elem) when length(list) == 0, do: false

  defp binary_search(list, elem) do
    binary_search(list, 0, length(list) - 1, elem)
  end

  defp binary_search(_list, left, right, _elem) when left > right, do: false

  defp binary_search(list, left, right, elem) do
    mid = div(right + left, 2)

    Enum.at(list, mid)
    |> binary_search(list, left, right, elem, mid)
  end

  defp binary_search(res, _list, _left, _right, e, _mid) when res == e, do: true

  defp binary_search(res, list, _left, right, e, mid) when res < e do
    binary_search(list, mid + 1, right, e)
  end

  defp binary_search(res, list, left, _right, e, mid) when res > e do
    binary_search(list, left, mid - 1, e)
  end

  # liniary search
  defp linear_search(list, k) do
    linear_search(list, 0, length(list) - 1, k)
  end

  defp linear_search(_list, left, right, _k) when left == right, do: false

  defp linear_search(list, left, right, k) do
    (Enum.at(list, left) + Enum.at(list, right))
    |> linear_search(list, left, right, k)
  end

  defp linear_search(res, _list, _left, _right, k) when res == k, do: true

  defp linear_search(res, list, left, right, k) when res > k do
    linear_search(list, left, right - 1, k)
  end

  defp linear_search(res, list, left, right, k) when res < k do
    linear_search(list, left + 1, right, k)
  end
end

[1, 2, 3, 4, 5]
|> SumTwoNumbersInList.is_two_numbers_exist?(10)
|> IO.inspect()
