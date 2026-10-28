package acceptance
import ("os";"testing")
func TestSyntheticFixture(t *testing.T){b,e:=os.ReadFile("fixture.txt");if e!=nil{t.Fatal(e)};if string(b)!="Atom remote acceptance fixture. Temporary branch; never merge.\n"{t.Fatal("fixture changed")}}
