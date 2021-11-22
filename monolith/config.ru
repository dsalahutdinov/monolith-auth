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

    puts result.inspect

    case req.path_info
    when /auth/
      if req.fetch_header('HTTP_AUTHORIZATION') == 'Token=123'
        [200, {'X-Auth-Identity' => '123'}, []]
      else
        [403, {"Content-Type" => "text/plain"}, []]
      end
    else
      if req.fetch_header("HTTP_X_AUTH_IDENTITY") { nil } == "123"
        [200, {"Content-Type" => "text/plain"}, [result]]
      else
        [403, {"Content-Type" => "text/plain"}, []]
      end
    end
  end
end

run Application.new
