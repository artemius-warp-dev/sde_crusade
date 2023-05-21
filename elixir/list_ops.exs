defmodule ListOps do
  # Please don't use any external modules (especially List or Enum) in your
  # implementation. The point of this exercise is to create these basic
  # functions yourself. You may use basic Kernel functions (like `Kernel.+/2`
  # for adding numbers), but please do not use Kernel functions for Lists like
  # `++`, `--`, `hd`, `tl`, `in`, and `length`.

  @spec count(list) :: non_neg_integer
  def count([]), do: 0
  def count(l), do: count(l, 0)
  defp count([], acc), do: acc
  defp count([_h | t], acc), do: count(t, acc + 1)

  @spec reverse(list) :: list
  def reverse(l) do
    reverse(l, [])
  end
  defp reverse([], acc), do: acc
  defp reverse([h | t], acc), do: reverse(t, [h | acc])

  @spec map(list, (any -> any)) :: list
  def map(l, f) do
    map(l, f, [])
  end

  defp map([], _, acc), do: acc |> reverse()
  defp map([h | t], f, acc), do: map(t, f, [f.(h) | acc])

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, f) do
    filter(l, f, [])
  end
  defp filter([], _f, acc), do: reverse(acc)
  defp filter([h | t], f, acc) do
    cond do
      f.(h) -> filter(t, f, [h | acc])
      true -> filter(t, f, acc)
    end
  end

  @type acc :: any
  @spec foldl(list, acc, (any, acc -> acc)) :: acc
  def foldl([], acc, _f), do: acc
  def foldl([h | t], acc, f), do: foldl(t, f.(h, acc), f)
    
  @spec foldr(list, acc, (any, acc -> acc)) :: acc
  def foldr(l, acc, f) do
    foldl(reverse(l), acc, f)
  end

  @spec append(list, list) :: list
  def append(a, b) do
    do_append(b, reverse(a))
  end
  defp do_append([], acc), do: reverse(acc)
  defp do_append([h | t], acc), do: do_append(t, [h | acc])

  @spec concat([[any]]) :: [any]
  def concat(ll) do
    concat(ll, [])
    |> reverse()
  end

  defp concat([], acc), do: acc
  defp concat([h | t], acc) when is_list(h) do
    concat(t, concat_nested(h, acc))
  end
  defp concat_nested([], acc), do: acc
  defp concat_nested([h | t], acc), do: concat_nested(t, [h | acc])
end
