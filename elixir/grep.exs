defmodule Grep do
  @spec grep(String.t(), [String.t()], [String.t()]) :: String.t()
  def grep(pattern, flags, files) do
      

      regex = 
        pattern
        |> append_x("-x" in flags)
        |> append_i("-i" in flags)
      files
      |> Enum.map(&read_in_stream(&1, regex, flags, length(files)))
      |> Enum.join()
  end

  defp append_x(pattern, x?) when x? do
    "^#{pattern}$" |> Regex.compile!()
  end

  defp append_x(pattern, _) when is_binary(pattern) do
    pattern |> Regex.compile!()
  end

  defp append_x(regex, _), do: regex

  defp append_i(pattern, i?) when is_binary(pattern) and i? do
    pattern |> Regex.compile!("i")
  end

  defp append_i(regex, i?) when i? do
    regex
    |> Regex.source()
    |> Regex.compile!("i")
  end
  
  defp append_i(pattern, _) when is_binary(pattern) do
    pattern |> Regex.compile!()
  end

  defp append_i(regex, _), do: regex
   

 

  defp read_in_stream(file, regex, flags, amount_of_files) do
    File.stream!(file)
    |> Stream.with_index(1)
    |> Stream.filter(fn {x, _c} ->
       match?(x, regex, "-v" in flags)
    end)
    |> Enum.to_list()
    |> render_stream("-n" in flags)
    |> add_filename_to_lines(file, amount_of_files)
    |> render_filenames(file, "-l" in flags)
    
  end

  defp match?(line, regex, v?) when v? do
    not String.match?(line, regex) 
  end
  defp match?(line, regex, _), do: String.match?(line, regex) 

  defp render_filenames(strs, file, l?) when l? and length(strs) > 0 do
    "#{file}\n"
  end

  defp render_filenames(strs, _, _) do
    strs
  end

  defp render_stream(list, n?) when n? do
    list
    |> Enum.map(fn {line, c} -> "#{c}:#{line}" end)
  end

  defp render_stream(list, _) do 
    list
    |> Enum.map(fn {line, _c} -> "#{line}" end)
  end

  defp add_filename_to_lines(lines, file, amount) when amount > 1 do
    lines
    |> Enum.map(fn line -> "#{file}:#{line}" end)
  end
  defp add_filename_to_lines(lines, _file, _amount), do: lines

end
