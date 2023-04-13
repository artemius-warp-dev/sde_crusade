defmodule SpaceAge do
  @type planet ::
          :mercury
          | :venus
          | :earth
          | :mars
          | :jupiter
          | :saturn
          | :uranus
          | :neptune

  @earth_year_seconds 31_557_600
  @planet_years %{
    mercury: 0.2408467,
    venus: 0.61519726,
    earth: 1.0,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }
  @doc """
  Return the number of years a person that has lived for 'seconds' seconds is
  aged on 'planet', or an error if 'planet' is not a planet.
  """
  @spec age_on(planet, pos_integer) :: {:ok, float} | {:error, String.t()}
  def age_on(planet, seconds) do
    seconds
    |> to_earth_years()
    |> to_planet_years(@planet_years[planet])
  end

  defp to_earth_years(seconds) do
    seconds / @earth_year_seconds
  end

  defp to_planet_years(_, nil), do: {:error, "not a planet"}

  defp to_planet_years(earth_years, planet) do
    {:ok, earth_years / planet}
  end
end

[:mercury, :venus, :earth, :mars, :jupiter, :saturn, :uranus, :neptune]
|> Enum.map(fn x ->
  {:ok, age} = SpaceAge.age_on(x, 31_557_600 * 25)
  {x, age}
end)
|> IO.inspect()
