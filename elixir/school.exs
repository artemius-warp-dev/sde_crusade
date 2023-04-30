defmodule School do
  @moduledoc """
  Simulate students in a school.

  Each student is in a grade.
  """

  @type school :: any()

  
  
  @doc """
  Create a new, empty school.
  """
  @spec new() :: school
  def new() do
    %{}
  end

  @doc """
  Add a student to a particular grade in school.
  """
  @spec add(school, String.t(), integer) :: {:ok | :error, school}
  def add(school, name, grade) do
    if name in roster(school) do
      {:error, school}
    else
    res = Map.update(school, grade, [name], fn names -> [name | names] end)
    {:ok, res}
    end
  end

  @doc """
  Return the names of the students in a particular grade, sorted alphabetically.
  """
  @spec grade(school, integer) :: [String.t()]
  def grade(school, grade) do
    Map.get(school, grade, []) |> Enum.sort()
  end

  @doc """
  Return the names of all the students in the school sorted by grade and name.
  """
  @spec roster(school) :: [String.t()]
  def roster(school) do
    school
    |> Enum.sort_by(fn {grade, _names} -> {grade} end)
    |> Enum.reduce([], fn {k, _v}, acc ->  acc ++ grade(school, k) end)
  end
end
