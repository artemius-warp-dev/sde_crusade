defmodule VariableLengthQuantity do
  import Bitwise

  @doc """
  Encode integers into a bitstring of VLQ encoded bytes
  """
  @spec encode(integers :: [integer]) :: binary
  def encode(integers) do
    Enum.reduce(integers, <<>>, &encode/2)
  end

  defp encode(0, acc), do: acc <> <<0>>
  defp encode(integer, acc), do: acc <> encode_integer(integer, 0, <<>>)

  defp encode_integer(0, _, acc), do: acc

  defp encode_integer(integer, msb, acc),
    do: encode_integer(integer >>> 7, 1, <<msb::1, integer::7, acc::binary>>)

  @doc """
  Decode a bitstring of VLQ encoded bytes into a series of integers
  """
  @spec decode(bytes :: binary) :: {:ok, [integer]} | {:error, String.t()}
  def decode(bytes) do
    decode(bytes, 0, [])
  end

  defp decode("", _, acc), do: {:ok, acc |> Enum.reverse()}
  defp decode(<<1::1, _x::7, "">>, _, _), do: {:error, "incomplete sequence"}

  defp decode(<<1::1, x::7, rest::bitstring>>, integer, acc),
    do: decode(rest, (integer <<< 7) + x, acc)

  defp decode(<<0::1, x::7, rest::bitstring>>, integer, acc),
    do: decode(rest, 0, [(integer <<< 7) + x | acc])
end
# encoded = <<
#   0xC0,
#   0x0,
#   0xC8,
#   0xE8,
#   0x56,
#   0xFF,
#   0xFF,
#   0xFF,
#   0x7F,
#   0x0,
#   0xFF,
#   0x7F,
#   0x81,
#   0x80,
#   0x0
# >>

#encoded = <<0xFF>>

#expected = {:ok, [0x2000, 0x123456, 0xFFFFFFF, 0x0, 0x3FFF, 0x4000]}
VariableLengthQuantity.encode([0]) |> IO.inspect()
#VariableLengthQuantity.decode(encoded) |> IO.inspect()
