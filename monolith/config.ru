require 'rack'
require 'json'
require_relative 'data'

class Application
  def call(env)
    req = Rack::Request.new(env)

    authorization = req.fetch_header("HTTP_AUTHORIZATION") { nil }
    current_user = if authorization != nil
      token = /Token=(?<token>\w*)/.match(authorization)['token']
      USER_BY_TOKEN[token] if token != ''
    elsif
      x_auth_identity = req.fetch_header("HTTP_X_AUTH_IDENTITY") { nil }
      USER_BY_ID[x_auth_identity]
    end
    

    case req.path_info
    when /products/
      json = JSON.generate(PRODUCTS)
      [200, {'Content-Type' => 'application/json'}, [json]]
    when /favorites/
      if current_user == nil
        [403, {"Content-Type" => "text/plain"}, []]
      else
        json = JSON.generate(FAVORITES_BY_USER_ID[current_user.id.to_s])
        [200, {"Content-Type" => "text/plain"}, [json]]
      end
    when /auth/
      if current_user != nil
        [200, {'X-Auth-Identity' => current_user.id.to_s}, []]
      else
        [403, {"Content-Type" => "text/plain"}, []]
      end
    end
  end
end

run Application.new
