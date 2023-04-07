defmodule ArmstrongNumber do
  @moduledoc """
  Provides a way to validate whether or not a number is an Armstrong number
  """

  @spec valid?(integer) :: boolean
  def valid?(number) do
    number
    |> to_string()
    |> String.graphemes()
    |> to_calculate()
    |> Kernel.==(number)
  end

  defp to_calculate(enum) do
    length = length(enum)

    enum
    |> Enum.map(&String.to_integer/1)
    |> Enum.reduce(0, fn x, acc ->
      x ** length + acc
    end)
  end
end

ArmstrongNumber.valid?(186_709_961_001_538_790_100_634_132_976_990) |> IO.inspect()
