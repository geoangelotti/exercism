datatype planet = Mercury | Venus | Earth | Mars
                | Jupiter | Saturn | Uranus | Neptune

fun age_on planet seconds =
    let
        val year = 365.25 * 24.0 * 60.0 * 60.0;
    in
        case planet of 
            Mercury => Real.fromInt seconds / 0.2408467 / year
            | Venus => Real.fromInt seconds / 0.61519726 / year
            | Earth => Real.fromInt seconds / 1.0 / year
            | Mars => Real.fromInt seconds / 1.8808158 / year
            | Jupiter => Real.fromInt seconds / 11.862615 / year
            | Saturn => Real.fromInt seconds / 29.447498 / year
            | Uranus => Real.fromInt seconds / 84.016846 / year
            | Neptune => Real.fromInt seconds / 164.79132 / year
    end

