{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Redundant id" #-}
import Prelude hiding (id, last)
import Data.Data (Typeable)
myfunc :: (Eq a, Num a) => a -> a
myfunc x
    | x == 0 = 0
    | otherwise = x + myfunc (x - 1)

lambda :: (HelloWorld -> Int) -> (Int -> Int)
lambda x = x

type HelloWorld = Int

class Myclass a where
    last :: [a] -> Bool


instance Myclass Int where
    last (x:xs)
        | null xs = False
        | x < 10 = True
        | otherwise = last xs

ans :: Int
ans = myfunc 10

newAns :: Int
newAns = myfunc 20

data Option a = Some a | None deriving(Eq, Typeable)


data JSON = JSON {
    name :: String,
    id :: Int
}

class Jsonify a where 
    toItem :: a -> (a -> (String, Int)) -> (String, Int)

instance Jsonify JSON where
    toItem :: JSON -> (JSON -> (String, Int)) -> (String, Int)
    toItem a f = f a


jsonDecode :: JSON -> (String, Int)
jsonDecode x = (name x, id x)

optionToInt :: Num z => Option z -> z
optionToInt a = case a of 
                    None -> 0
                    Some z -> z + 10



                    