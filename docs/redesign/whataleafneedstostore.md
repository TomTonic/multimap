Diese Datei enthält Überlegungen dazu, was eine Seite wenigstens wissen/speichern muss.

Eine Seite habe die verfügbare Kapazität B.
Ein Schlüssel k habe die Rest-Länge R; die Werte K1, K2, K3 usw. bezeichnen die Längen der Schlüssel k1, k2, k3 usw.
Ein Wert w zu einem Schlüssek k habe die Länge Wk; die Werte W1k, W2k, W3k usw. bezeichnen die Längen von mehreren Werten w1, w2, w3 usw. zum Schlüssek k; die Werte Wk1, Wk2, Wk3 usw. bezeichnen die Längen der (eindeutigen) Werte w zum den Schlüsseln k1, k2, k3; die Werte W1k1, W2k1, W1k2, W2k2 usw. bezeichnen die Längen der Werte w1 und w2 zu k1 und w1' und w2' zu k2.

Alle Werte sind immer in byte.

# Fall 1: Ein Schlüssel, ein Wert

## Fall 1.a: R + W <= B

Die Seite speichert den kompletten Schlüssel und den kompletten Wert

## Fall 1.b: R + W > B

### Fall 1.b.i: R < B und W < B

Wenn möglich und wenn das reicht: die Seite wird vergrößert.
Sonst: Die Seite speichert den kompletten Schlüssel und eine Referenz auf den Wert

### Fall 1.b.ii: R > B und W < B

Wenn möglich und wenn das reicht: die Seite wird vergrößert.
Sonst: Die Seite speichert eine Referenz auf den Schlüssel (Array oder String) und den direkt eingebetteten Wert

### Fall 1.b.iii: R < B und W > B

Wenn möglich und wenn das reicht: die Seite wird vergrößert.
Sonst: Die Seite speichert den kompletten Schlüssel und eine Referenz auf den Wert

### Fall 1.b.iv: R > B und W > B

Wenn möglich und wenn das reicht: die Seite wird vergrößert.
Sonst: Die Seite speichert eine Referenz auf den Schlüssel (Array oder String) und eine Referenz auf den Wert

# Fall 2: Ein Schlüssel, mehrere Werte

## Fall 2.a: R + W1k + ... + Wnk <= B

Die Seite speichert den kompletten Schlüssel und die kompletten Werte

## Fall 2.b: R + W1k + ... + Wnk > B

## Fall 2.b.i: R < B und W1k + ... + Wnk > B-R

Wenn möglich und wenn das reicht: die Seite wird vergrößert.
Sonst: Die Seite speichert den kompletten Schlüssel und eine Referenz auf ein Set3/Array/vset

## Fall 2.b.i: R > B und W1k + ... + Wnk < B

Die Seite speichert eine Referenz auf den Schlüssel (Array oder String) und die direkt eingebetteten Werte

## Fall 2.b.i: R < B und W1k + ... + Wnk > B-R

# Fall 3: Mehrere Schlüssel, je ein Wert

...

# Fall 4: Mehrere Schlüssel, mindestens einer mit mehreren Werten

... Stronzenglubber ...

Seiten haben immer 128 bytes, 256 bytes, 384 bytes oder 512 bytes
Seiten werden nur verkleinert, wenn das 50% Speicher spart
Frei gewordene Seiten werden in einem LIFO Page-cache gespeichert

# Fall 1: SKSV - Single Key Single Value

Wird als Spezialfall von SKMV behandelt

# Fall 2: SKMV - Single Key Multiple Values

## Fall 2a: Alles passt auf eine Seite (Prio 1)

TB LR V1 V2 V3 V4 V5 V6 | k1 k2 k3 ... kLR | v11 v12 v13 ... v1V1 | v21 v22 v23 ... v2V2 | ...
0                       8                  8+LR                   8+LR+V1                8+LR+V1+V2

TB - TypeByte mit höchstwertigem Bit an LR ausgeliehen (als 2 Byte-Werte für diesen Typ)
LR - Länge des Schlüssel-Restes in Bytes (8 bit + 1 bit aus TB)
V1 - Länge des Value 1 in Byte (nur 8 bit, 0 wenn kein Wert vorhanden)
V2 - Länge des Value 2 in Byte (nur 8 bit, 0 wenn kein weiterer Wert vorhanden)
...
k1, k2, k3 - Byte-Werte des Schlüssel-Restes
v1x - Byte-Werte des ersten Wertes
v2x - Byte-Werte des zweiten Wertes
...

* Schlüsselreste bis 503 bytes
* bis zu 6 Werte mit bis zu 256 bytes

## Fall 2b: Der Schlüssel-Rest passt Seite, die Werte passen aber nicht zusätzlich drauf (Prio 2)

TB LR _ _ _ _ _ _ | Pv | k1 k2 k3 ... kLR | _
0                 8    16                 16+LR

TB - TypeByte mit höchstwertigem Bit an LR ausgeliehen (als 2 Byte-Werte für diesen Typ)
LR - Länge des Schlüssel-Restes in Bytes (8 bit + 1 bit aus TB)
V1 - Länge des Value 1 in Byte (nur 8 bit, 0 wenn kein Wert vorhanden)
V2 - Länge des Value 2 in Byte (nur 8 bit, 0 wenn kein weiterer Wert vorhanden)
...
Pv - Pointer auf das vset
k1, k2, k3 - Byte-Werte des Schlüssel-Restes

* Schlüsselreste bis 496 bytes
* beliebig viele Werte, ggf. direkt typisierter Zeiger auf Array oder Set3

## Fall 2c: Der Schlüssel-Rest ist zu lang für die Seite, die Werte passen aber noch drauf  (Prio 3)

TB V1 V2 V3 V4 V5 V6 V7 | Hk | Pk | v11 v12 v13 ... v1V1 | v21 v22 v23 ... v2V2 | ...
0                       8    16   24                     24+V1                  16+V1+V2

TB - TypeByte
V1 - Länge des Value 1 in Byte (nur 8 bit, 0 wenn kein Wert vorhanden)
V2 - Länge des Value 2 in Byte (nur 8 bit, 0 wenn kein weiterer Wert vorhanden)
...
Hk - Hashwert des gesamten Schlüssels (?)
Pk - Pointer auf den kompletten Schlüsselrest
k1, k2, k3 - Byte-Werte des Schlüssel-Restes
v1x - Byte-Werte des ersten Wertes
v2x - Byte-Werte des zweiten Wertes
...

* Schlüsselreste über 503 bytes, beliebig lang
* bis zu 7 Werte mit bis zu 256 bytes

## Fall 2d: Der Schlüssel-Rest passt nicht auf die Seite und die Werte passen auch nicht zusätzlich drauf (Prio 4)

TB _ _ _ _ _ _ _ | Pk | Pv |
0                 8    16   24

TB - TypeByte mit höchstwertigem Bit an LR ausgeliehen (als 2 Byte-Werte für diesen Typ)
LR - Länge des Schlüssel-Restes in Bytes (8 bit + 1 bit aus TB)
V1 - Länge des Value 1 in Byte (nur 8 bit, 0 wenn kein Wert vorhanden)
V2 - Länge des Value 2 in Byte (nur 8 bit, 0 wenn kein weiterer Wert vorhanden)
...
Pk - Pointer auf den Key
Pv - Pointer auf das vset
k1, k2, k3 - Byte-Werte des Schlüssel-Restes

* Schlüsselreste bis 496 bytes
* beliebig viele Werte, ggf. direkt typisierter Zeiger auf Array oder Set3

# Fall 3: MKSV - Multiple Key Single Values

TB CP R1 R2 R3 V1 V2 V3 | cp1 cp2 ... cpCP | k11 k12 k13 ... k1R1 | k21 k22 k23 ... k2R2 | k31 k32 k33 ... k3R3 | v11 v12 v13 ... v1V1 | ...
0                       8                  8+CP                   8+CP+R1                8+CP+R1+R2             8+CP+R1+R2+R3


TB CP R1 V1 R2 V2 R3 V3 | V4 R4 V7 R7 V7 R7 V7 R7 | cp1 cp2 ... cpCP | r11 r21 r31 ... rR11 | r11 r12 r13 ... r1R1 | v11 v12 v13 ... v1V1 | ...
0                       8                         16+CP               8+CP+R1               8+CP+L1+V1                8+LR+V1+V2

CP - Länge des Common Prefix des Restschlüssels
cp1, cp2, cp3 - Bytes des Common Prefix
Rx - Länge des Restschlüssel x
k11, k12, k13 - Bytes des Restschlüssel k1

* Common Prefix bis 256 bytes
* Schlüsselreste (ohne prefix) bis 255 bytes (0 heißt kein weiterer Schlüssel vorhanden)
* Werte bis zu 256 bytes
* Entweder bis zu 3 Schlüssel und Werte oder bis zu 7 Schlüssel und Werte

# Fall 4: MKMV - Multiple Key Multiple Values

Sobald ein Schlüssel in MKSV einen weiteren Wert eingefügt bekommt, wird er aus der Seite entfernt und ein SKMV für ihn angelegt