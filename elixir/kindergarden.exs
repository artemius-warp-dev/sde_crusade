defmodule Garden do
  @doc """
    Accepts a string representing the arrangement of cups on a windowsill and a
    list with names of students in the class. The student names list does not
    have to be in alphabetical order.

    It decodes that string into the various gardens for each student and returns
    that information in a map.
  """
  
  @children ~w(alice bob charlie david eve fred ginny harriet ileana joseph kincaid larry)a

  @spec info(String.t(), list) :: map
  def info(info_string, student_names \\ @children) do
    flowers = 
    info_string
      |> String.split("\n")
      |> Enum.map(&(String.graphemes(&1) |> match() |> Enum.chunk_every(2)))
      |> Enum.zip()
      |> Enum.map(fn {f1, f2} -> (f1 ++ f2) |> List.to_tuple() end)

    map = 
        student_names
        |> Enum.sort()
        |> Enum.zip(flowers)
        |> Map.new()

    @children ++ student_names 
    |> Enum.reduce(map, fn key, acc -> Map.put_new(acc, key, {}) end)
    
  end

  defp match(strs) when is_list(strs) do
    strs
    |> Enum.map(&match/1)
  end
  
  defp match(str) do
    case str do
      "G" -> :grass
      "C" -> :clover
      "R" -> :radishes
      "V" -> :violets
    end
  end
end
