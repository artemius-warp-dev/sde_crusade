defmodule IsbnVerifier do
  @doc """
    Checks if a string is a valid ISBN-10 identifier

    ## Examples

      iex> IsbnVerifier.isbn?("3-598-21507-X")
      true

      iex> IsbnVerifier.isbn?("3-598-2K507-0")
      false

  """
  @spec isbn?(String.t()) :: boolean
  def isbn?(isbn) do
    if isbn =~ ~r/^(\d-?){9}(\d|X)$/ do
      Regex.scan(~r/[0-9]|X$/, isbn)
      |> List.flatten()
      |> Enum.map(&to_str/1)
      |> Enum.reduce({0, 10}, fn
        x, {acc, count} -> {x * count + acc, count - 1}
      end)
      |> then(fn {res, _} -> rem(res, 11) == 0 end)
    end
  end

  defp to_str("X"), do: 10
  defp to_str(ch), do: String.to_integer(ch)
end
