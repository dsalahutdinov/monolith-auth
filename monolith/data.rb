require "ostruct"

USER_BY_ID = {
  '123' => OpenStruct.new(id: 123),
  '234' => OpenStruct.new(id: 234),
  '345' => OpenStruct.new(id: 345)
}

USER_BY_TOKEN = {
  'qwerty' => OpenStruct.new(id: 123),
  'asdfgh' => OpenStruct.new(id: 234),
  'zxcvbn' => OpenStruct.new(id: 345)
}

PRODUCTS = [
  { id: '1', name: 'Potato' },
  { id: '2', name: 'Tomato' },
  { id: '3', name: 'Onion' },
  { id: '4', name: 'Carrot' },
  { id: '5', name: 'Canary' }
]

FAVORITES_BY_USER_ID = {
  '123' => [{product_id: '1'}, {product_id: '2'}],
  '234' => [{product_id: '1'}, {product_id: '3'}],
  '345' => [{product_id: '1'}, {product_id: '4'}],
}
