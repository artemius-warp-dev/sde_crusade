defmodule SecretHandshake do
  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.

  If the following bits are set, include the corresponding action in your list
  of commands, in order from lowest to highest.

  1 = wink
  10 = double blink
  100 = close your eyes
  1000 = jump

  10000 = Reverse the order of the operations in the secret handshake
  """
  @spec commands(code :: integer) :: list(String.t())
  def commands(code) do
    []
    |> do_shift(code)
  end
  
  defp do_shift(acc, code) do
    do_shift(acc, code, 1)
  end
  defp do_shift(acc, _code, 6), do: acc |> Enum.reverse()
  defp do_shift(acc, code, pos) do
    res = Bitwise.>>>(code, pos-1) |> Bitwise.&&&(1) == 1
    cond do
      res and pos == 1 -> ["wink" | acc]
      res and pos == 2 -> ["double blink" | acc]
      res and pos == 3 -> ["close your eyes" | acc]
      res and pos == 4 -> ["jump" | acc]
      res and pos == 5 -> Enum.reverse(acc)
      true -> acc
    end
    |> do_shift(code, pos+1)
  end
end
