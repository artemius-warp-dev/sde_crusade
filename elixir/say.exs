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
    40 => "forty",
    50 => "fifty",
    60 => "sixty",
    70 => "seventy",
    80 => "eighty",
    90 => "ninety",
    100 => "hundred",
    1000 => "thousand",
    1_000_000 => "million",
    1_000_000_000 => "billion"
  }

  @doc """
  Translate a positive integer into English.
  """
  @spec in_english(integer) :: {atom, String.t()}
  def in_english(number) when number in 0..999_999_999_999 do
    split(number)
    #|> IO.inspect()
    |> combine()
    |> then(fn res -> {:ok, res} end)
  end

  def in_english(_), do: {:error, "number is out of range"}

  defp split(n) when n < 100, do: {div(n, 10) * 10, rem(n, 10)}
  defp split(n) when n < 1000 and rem(n, 100) == 0, do: {div(n, 100), 100}
  defp split(n) when n < 1000, do: {div(n, 100), 100, rem(n, 100)}
  defp split(n) when n < 10_000 and rem(n, 1000) == 0, do: {div(n, 1000), 1000}
  defp split(n) when n < 10_000, do: {div(n, 1000), 1000, rem(n, 1000)}


  defp split(n) when n < 100_000 and rem(n, 10_000) == 0, do: {div(n, 10_000), 10_000}
  defp split(n) when n < 100_000, do: {div(n, 10_000) * 10, rem(n, 10_000)}

  defp split(n) when n < 1_000_000 and rem(n, 100_000) == 0, do: {div(n, 100_000), 100_000}
  defp split(n) when n < 1_000_000, do: {div(n, 100_000), 100, rem(n, 100_000)}


  defp split(n) when n < 10_000_000 and rem(n, 1_000_000) == 0, do: {div(n, 1_000_000), 1_000_000}
  defp split(n) when n < 10_000_000, do: {div(n, 1_000_000), 1_000_000, rem(n, 1_000_000)}

  defp split(n) when n < 100_000_000 and rem(n, 10_000_00) == 0,
    do: {div(n, 10_000_000), 10_000_000}

  defp split(n) when n < 100_000_000 do
    {div(n, 10_000_000) * 10, rem(n, 10_000_000)}
  end

  defp split(n) when n < 1_000_000_000 and rem(n, 100_000_000) == 0,
    do: {div(n, 100_000_000), 100_000_000}



  defp split(n) when n < 1_000_000_000,
    do: {div(n, 100_000_000), 100, rem(n, 100_000_000)}


  defp split(n) when n < 10_000_000_000 and rem(n, 1_000_000_000) == 0 do: {div(n, 1_000_000_000), 1_000_000_000}

  defp split(n) when n < 10_000_000_000, do: {div(n, 1_000_000_000), 1_000_000_000, rem(n, 1_000_000_000)}



  defp split(n) when n < 100_000_000_000 and rem(n, 10_000_000_000) == 0,
    do: {div(n, 10_000_000_000), 10_000_000_000}

  defp split(n) when n < 100_000_000_000,
    do: {div(n, 10_000_000_000) * 10, rem(n, 10_000_000_000)}

  defp split(n) when n < 1000_000_000_000,
    do: {div(n, 100_000_000_000), 100, rem(n, 100_000_000_000)}

  defp combine({0, units}), do: units(units)
  defp combine({dozens, 0}), do: dozens(dozens)
  defp combine({10, units}), do: "#{units(units)}#{dozens(10)}"
  defp combine({unit, dozen}) when unit < 10 do
    #IO.inspect({unit, dozen})
    "#{units(unit)} #{dozens(dozen)}"
  end

  defp combine({unit, dozen, reminder}) do
    "#{units(unit)} #{dozens(dozen)} #{combine(split(reminder))}"
  end

  defp combine({dozens, units}) when units > 10 do
    "#{dozens(dozens)}-#{combine(split(units))}"
  end

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

Say.in_english(987_654_321_123) |> IO.inspect()
#"nine hundred eighty-seven billion six hundred fifty-four million three hundred twenty-one thousand one hundred twenty-three"}
