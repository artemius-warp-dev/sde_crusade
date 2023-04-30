defmodule Proverb do
  @doc """
  Generate a proverb from a list of strings.
  """
  @spec recite(strings :: [String.t()]) :: String.t()
  def recite(strings) do
    strings
    |> proverbs({"", []})
    |> then(fn
        [h | t] -> t ++ [h] 
        acc -> acc
    end)
    |> Enum.join("\n")
  end

 
  defp proverbs([], {_,acc}), do: Enum.reverse(acc)
  defp proverbs([h | t], {"", []}), do: proverbs(t, {h, ["And all for the want of a #{h}.\n"]})
  defp proverbs([h | t], {prev, acc}), do: proverbs(t, {h, ["For want of a #{prev} the #{h} was lost." | acc]})
end
