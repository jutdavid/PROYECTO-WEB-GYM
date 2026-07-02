

type CicloEntrenamiento (
	ID          INTEGER PRIMARY KEY,
	AtletaID    INTEGER NOT NULL,
	Estado      TEXT    NOT NULL,
	FechaInicio TEXT    NOT NULL
);

type Microciclo (
	ID                   INTEGER PRIMARY KEY,
	CicloEntrenamientoID INTEGER NOT NULL,
	NumeroSemana         INTEGER NOT NULL,
	EnfoqueEspecifico    TEXT    NOT NULL,
	FechaInicio          INTEGER NOT NULL
);

type EvaluacionCiclo (
	ID                   INTEGER PRIMARY KEY,
	CicloEntrenamientoID INTEGER NOT NULL,
	NivelFatiga          INTEGER NOT NULL,
	Comentarios          TEXT    NOT NULL,
	FechaEvaluacion      INTEGER NOT NULL
);