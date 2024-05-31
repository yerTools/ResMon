-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Graph.Mutation exposing (..)

import Api.ScalarCodecs
import Graph.InputObject
import Graph.Interface
import Graph.Object
import Graph.Scalar
import Graph.Union
import Graphql.Internal.Builder.Argument as Argument exposing (Argument)
import Graphql.Internal.Builder.Object as Object
import Graphql.Internal.Encode as Encode exposing (Value)
import Graphql.Operation exposing (RootMutation, RootQuery, RootSubscription)
import Graphql.OptionalArgument exposing (OptionalArgument(..))
import Graphql.SelectionSet exposing (SelectionSet)
import Json.Decode as Decode exposing (Decoder)


workClock :
    SelectionSet decodesTo Graph.Object.WorkClockMutation
    -> SelectionSet decodesTo RootMutation
workClock object____ =
    Object.selectionForCompositeField "workClock" [] object____ Basics.identity