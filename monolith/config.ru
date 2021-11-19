require 'rack'


class Application
  def call(env)
    req = Rack::Request.new(env)
    http_headers = req
      .each_header
      .select { |k, v|  k.start_with? 'HTTP_'}
      .collect {|key, val| [key.sub(/^HTTP_/, ''), val]}
      .collect {|key, val| "#{key}: #{val}"}

    result = (["PATH: #{req.path_info}", "\n"] + http_headers).join("\n")

    case req.path_info
    when /auth/
      [200, {"Content-Type" => "text/plain"}, [result]]
    else
      [200, {"Content-Type" => "text/plain"}, [result]]
    end
  end
end

run Application.new
