defmodule AllYourBase do
  @doc """
  Given a number in input base, represented as a sequence of digits, converts it to output base,
  or returns an error tuple if either of the bases are less than 2
  """

  @spec convert(list, integer, integer) :: {:ok, list} | {:error, String.t()}
  def convert(digits, input_base, output_base)

  def convert(digits, input_base, output_base) when output_base < 2,
    do: {:error, "output base must be >= 2"}

  def convert(digits, input_base, output_base) when input_base < 2,
    do: {:error, "input base must be >= 2"}

  def convert(digits, input_base, output_base) do
    with {:ok, ^digits} <- check_digits(digits, input_base) do
      res =
        to_decimal_from(input_base, digits)
        |> from_decimal_to(output_base, [])

      {:ok, res}
    end
  end

  defp check_digits(digits, input_base) do
    case Enum.filter(digits, &(&1 < 0 or &1 >= input_base)) do
      [] -> {:ok, digits}
      _ -> {:error, "all digits must be >= 0 and < input base"}
    end
  end

  defp to_decimal_from(input_base, digits) do
    {number, _} =
      digits
      |> Enum.reverse()
      |> Enum.reduce({0, 0}, fn x, {acc, c} ->
        acc = (acc + x * :math.pow(input_base, c)) |> round()
        {acc, c + 1}
      end)

    number
  end

  defp from_decimal_to(number, output_base, acc) do
    case div(number, output_base) do
      0 -> [rem(number, output_base) | acc]
      res -> from_decimal_to(res, output_base, [rem(number, output_base) | acc])
    end
  end
end

AllYourBase.convert([1, -1, 1, 0, 1, 0], 2, 10) |> IO.inspect()
