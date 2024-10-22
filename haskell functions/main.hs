{-# OPTIONS_GHC -Wno-unrecognised-pragmas #-}
{-# HLINT ignore "Redundant id" #-}
{-# HLINT ignore "Use camelCase" #-}
import Prelude hiding (id, last)
import Data.Data (Typeable)
import Data.Void (Void)
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
    last :: [Int] -> Bool
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


c :: String
c = "aman"

index :: Int -> Int -> [Char] -> Char
index t l x = case x of 
        [] -> 'x'
        (y:x) -> if l == t then y else index (t+1) l x        

reversed :: [Int] -> [Int]
reversed [] = []
reversed (x:xs)
    | null xs = [x]
    | otherwise =  reversed xs ++ [x]

data Socket p i = Port p | TCP i | Socket { port :: p, tcp :: i } deriving(Show)

socket :: Socket Int Void
socket = Port 8000

equation :: String -> Int -> Socket Int String
equation x = Port


data Tree a = Node a (Tree a) | Leaf a deriving(Show)

get_tree :: Int -> Tree Int -> Tree Int
get_tree x y = case y of
                Leaf temp -> Node temp (Leaf 0)
                Node h (Node z k) -> Node h (Leaf z)
                _ -> Leaf 100  