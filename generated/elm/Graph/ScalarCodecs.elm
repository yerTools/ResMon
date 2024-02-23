-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Graph.ScalarCodecs exposing (..)

import Graph.Scalar exposing (defaultCodecs)
import Json.Decode as Decode exposing (Decoder)


type alias Any_ =
    Graph.Scalar.Any_


type alias FieldSet_ =
    Graph.Scalar.FieldSet_


type alias Id =
    Graph.Scalar.Id


codecs : Graph.Scalar.Codecs Any_ FieldSet_ Id
codecs =
    Graph.Scalar.defineCodecs
        { codecAny_ = defaultCodecs.codecAny_
        , codecFieldSet_ = defaultCodecs.codecFieldSet_
        , codecId = defaultCodecs.codecId
        }