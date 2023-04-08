defmodule RomanNumerals do
  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) when number < 4000 do
    {"", number}
    |> add_m() #CM
    |> add_d() #CD
    |> add_c() #XC
    |> add_l() #XL
    |> add_x() #IX
    |> add_v() #IV
    |> add_i()
    |> then(fn {res, _} -> res end)
  end


  
  defp add_m({res, number}) when number < 1000 and number >= 900, do: {"#{res}CM", number - 900}
  defp add_m({res, number}) when div(number, 1000) > 0, do: add(res, number , 1000, "M") |> add_m()
  defp add_m({res, number}), do: {res, number}

  defp add_d({res, number}) when number < 500 and number >= 400, do: {"#{res}CD", number - 400}
  defp add_d({res, number}) when div(number, 500) > 0, do: add(res, number , 500, "D") |> add_d()
  defp add_d({res, number}), do: {res, number}

  defp add_c({res, number}) when number < 100 and number >= 90, do: {"#{res}XC", number - 90}
  defp add_c({res, number}) when div(number, 100) > 0, do: add(res, number , 100, "C") |> add_c()
  defp add_c({res, number}), do: {res, number}

  defp add_l({res, number}) when number < 50 and number >= 40, do: {"#{res}XL", number - 40}
  defp add_l({res, number}) when div(number, 50) > 0, do: add(res, number , 50, "L") |> add_l()
  defp add_l({res, number}), do: {res, number}

  defp add_x({res, number}) when number < 10 and number >= 9, do: {"#{res}IX", number - 9}
  defp add_x({res, number}) when div(number, 10) > 0, do: add(res, number, 10, "X") |> add_x()
  defp add_x({res, number}), do: {res, number}

  defp add_v({res, number}) when number < 5 and number >= 4, do: {"#{res}IV", number - 4}
  defp add_v({res, number}) when div(number, 5) > 0, do: add(res, number, 5, "V") |> add_v()
  defp add_v({res, number}), do: {res, number}
    
  defp add_i({res, number}) when div(number, 1) > 0, do: add(res, number, 1, "I")
  defp add_i({res, number}), do: {res, number}

  defp add(res, number, divider, roman_digit) do
    amount_of_ch = div(number, divider)
    ch = Enum.reduce(0..amount_of_ch-1, "", fn _, acc -> "#{acc}#{roman_digit}" end)
    {"#{res}#{ch}", number - (amount_of_ch * divider)}
  end  

end
