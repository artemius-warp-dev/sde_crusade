defmodule RotationalCipher do
  @alphabet "abcdefghijklmnopqrstuvwxyz"
  @capital_alphabet @alphabet |> String.upcase()

  @doc """
  Given a plaintext and amount to shift by, return a rotated string.
  
  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
    |> String.graphemes()
    |> Enum.map(&(shift(&1, shift)))
    |> Enum.join("")
  end

  defp shift(ch, shift) do
    cond do
      String.contains?(@alphabet, ch) ->
        do_shift(@alphabet, ch, shift)
      String.contains?(@capital_alphabet, ch) ->
        do_shift(@capital_alphabet, ch, shift)
      true -> 
        ch
    end
  end

  defp do_shift(alphabet, ch, shift) do
    [h, tail] = String.split(alphabet, ch)
    "#{tail}#{h}#{ch}"
    |> to_charlist()
    |> Enum.at(shift - 1) 
    |> then(fn x -> [x] end)
  end
end



