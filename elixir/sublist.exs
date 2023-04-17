defmodule Sublist do
  @doc """
  Returns whether the first list is a sublist or a superlist of the second list
  and if not whether it is equal or unequal to the second list.
  """
  def compare(a, b) do
   cond do 
      a == b -> :equal
      sublist?(a,b) -> :sublist
      sublist?(b,a) -> :superlist
      true -> :unequal
   end
  end

  defp sublist?(a,b), do: sublist?(a,b, {a,b})
  defp sublist?([], _, _), do: true
  defp sublist?(_, [], _), do: false
  defp sublist?([h1 | t1] = p1, [h2 | t2] = p2, _) when h1 === h2, do: sublist?(t1, t2, {p1,p2})
  

  defp sublist?(_a, _, {p1, [_h | p2]}), do: sublist?(p1, p2, {p1,p2})
end
