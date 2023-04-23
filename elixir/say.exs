defmodule Say do

  @units %{
    1 => "one",
    2 => "two",
    3 => "three",
    4 => "four",
    5 => "five",
    6 => "six",
    7 => "seven",
    8 => "eight",
    9 => "nine",
    0 => "zero"
  }
  @dozens %{
    10 => "teen",
    20 => "twenty",
    30 => "thirty",
    40 => "fourty",
    50 => "fifty",
    60 => "sixty",
    70 => "seventy",
    80 => "eighty",
    90 => "ninety"
  }


  @doc """
  Translate a positive integer into English.
  """
  @spec in_english(integer) :: {atom, String.t()}
  def in_english(number) when number in 0..999_999_999_999 do
    split(number)
    |> IO.inspect()
    |> combine()
    |> then(fn res -> {:ok, res} end)
  end

  def in_english(_), do: {:error, "number is out of range"}

  #defp split(number) when number > 10, do: {div(number, 10) * 10, rem(number, 10)}
  #defp split(number) when number >= 100 do
  #  hundreds = div(number, 100) * 100
  #  {div(number, 100), div(number - hundreds, 10) * 10, }
  #end

  #defp split(number), do: {0, number}

  

  defp split(n) when n < 100, do: {}
  defp split(n) when n < 1000
  defp split(n) when n < 10_000
  defp split(n) when n < 100_000
  defp split(n) when n < 1000_000
  defp split(n) when n < 10_000_000
  defp split(n) when n < 100_000_000
  defp split(n) when n < 1_000_000_000
  defp split(n) when n < 10_000_000_000
  defp split(n) when n < 100_000_000_000
  defp split(n) when n < 1000_000_000_000, do: [div(n, 100_000_00)]


  defp combine({0, units}), do: units(units)
  defp combine({dozens, 0}), do: dozens(dozens)
  defp combine({10, units}), do: "#{units(units)}#{dozens(10)}"

  defp combine({dozens, units}) do
    "#{dozens(dozens)}-#{units(units)}"
  end

  defp units(unit) do
    @units[unit]
  end
  
  defp dozens(dozen) do
    @dozens[dozen]
  end

end
