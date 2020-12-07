package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bagAndCount struct {
	bag string
	count int
}

func main() {
	input := ReadInput()
	bagMap := map[string][]bagAndCount{}
	for _, l := range(input) {
		pars := strings.Split(l, " bags contain ")
		k := pars[0]
		vs := strings.Split(pars[1], ", ")
		vals := []bagAndCount{}
		for _, v := range(vs) {
			bac := strings.Split(v, " ")
			c, err := strconv.Atoi(bac[0])
			if err != nil {
				break
			}
			b := fmt.Sprintf("%s %s", bac[1], bac[2])
			vals = append(vals, bagAndCount{b, c})
		}
		bagMap[k] = vals
	}
	fmt.Printf("%d", contains("shiny gold", &bagMap))
}

func contains(bag string, bagMap *map[string][]bagAndCount) int {
	total := 0
	for _, bac := range((*bagMap)[bag]) {
		total += bac.count * (1 + contains(bac.bag, bagMap))
	}	
	return total
}

func ReadFakeInput() []string {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	return strings.Split(input, "\n")
}

func ReadInput() []string {
	input := `vibrant bronze bags contain 3 dim olive bags.
shiny teal bags contain 1 posh green bag, 5 pale indigo bags, 1 mirrored purple bag.
striped aqua bags contain 5 bright orange bags.
clear chartreuse bags contain 3 dotted black bags, 2 wavy olive bags.
light lime bags contain 1 posh silver bag, 5 clear orange bags, 2 light olive bags, 3 dull maroon bags.
light olive bags contain 4 striped turquoise bags.
shiny purple bags contain 2 posh silver bags, 3 striped silver bags, 5 shiny beige bags, 2 plaid chartreuse bags.
mirrored crimson bags contain 2 faded cyan bags.
shiny turquoise bags contain 5 dull purple bags.
dim red bags contain 2 dim salmon bags, 2 faded orange bags, 5 muted aqua bags.
vibrant yellow bags contain 5 mirrored white bags, 5 vibrant blue bags, 3 mirrored lavender bags, 1 wavy cyan bag.
posh salmon bags contain 1 dull black bag, 1 striped indigo bag, 1 muted silver bag, 2 vibrant crimson bags.
pale black bags contain 1 plaid cyan bag.
dotted salmon bags contain 3 wavy brown bags, 3 pale coral bags, 1 light maroon bag.
posh orange bags contain 5 muted green bags, 3 striped violet bags.
dull maroon bags contain 2 clear brown bags, 5 posh silver bags, 5 mirrored coral bags, 2 dim lavender bags.
bright lavender bags contain 3 dark chartreuse bags, 1 mirrored chartreuse bag, 2 striped orange bags, 4 striped bronze bags.
plaid white bags contain 2 pale aqua bags.
posh teal bags contain 1 muted crimson bag, 2 dark fuchsia bags, 2 dim black bags, 4 plaid cyan bags.
wavy maroon bags contain 2 dull magenta bags, 3 dark red bags, 5 dull green bags, 4 bright turquoise bags.
plaid teal bags contain 5 plaid plum bags, 3 light magenta bags.
plaid plum bags contain 3 striped lime bags, 5 clear maroon bags, 3 muted plum bags.
muted purple bags contain 5 muted fuchsia bags, 4 pale tomato bags.
dark gold bags contain 5 dim lime bags, 3 clear orange bags, 4 drab crimson bags, 1 faded cyan bag.
striped coral bags contain 4 pale aqua bags, 5 clear silver bags.
shiny chartreuse bags contain 1 muted plum bag, 3 vibrant tomato bags.
bright salmon bags contain 5 pale gold bags, 1 muted gold bag, 5 dark gray bags, 4 dull cyan bags.
dark crimson bags contain 1 plaid turquoise bag.
light coral bags contain 1 muted brown bag, 2 striped black bags, 5 dark gray bags.
dotted lavender bags contain 1 bright turquoise bag.
posh red bags contain 2 muted green bags.
dim turquoise bags contain 4 dull chartreuse bags.
posh lime bags contain 5 mirrored yellow bags, 1 striped silver bag.
wavy black bags contain 5 striped cyan bags, 4 wavy red bags, 2 dotted coral bags.
dotted brown bags contain 1 dim gray bag, 1 plaid tomato bag.
mirrored red bags contain 4 posh aqua bags, 4 dark gray bags, 5 dark turquoise bags.
plaid tan bags contain 4 plaid black bags, 4 dull fuchsia bags, 1 plaid plum bag, 3 dark chartreuse bags.
plaid bronze bags contain 2 muted lavender bags, 3 faded cyan bags, 3 mirrored chartreuse bags, 1 dull coral bag.
mirrored silver bags contain 3 dull bronze bags, 3 dim tomato bags.
shiny fuchsia bags contain 3 bright maroon bags, 1 vibrant tomato bag, 4 posh bronze bags, 1 striped bronze bag.
dotted plum bags contain 5 wavy fuchsia bags.
dim bronze bags contain 3 shiny red bags, 5 dotted chartreuse bags.
faded crimson bags contain 3 bright olive bags, 1 dark bronze bag, 5 drab crimson bags.
striped blue bags contain 4 drab blue bags.
posh purple bags contain 1 bright blue bag, 4 light black bags, 1 dotted violet bag.
shiny lavender bags contain 4 mirrored bronze bags.
dull beige bags contain 2 wavy chartreuse bags.
dim blue bags contain 1 bright magenta bag, 5 muted red bags.
pale yellow bags contain 1 dotted white bag.
mirrored blue bags contain 3 striped tan bags.
mirrored turquoise bags contain 2 plaid red bags, 5 muted red bags, 2 muted green bags.
faded fuchsia bags contain 3 wavy tomato bags, 1 vibrant red bag, 1 dotted green bag, 2 posh plum bags.
drab silver bags contain 2 muted fuchsia bags, 4 dotted gray bags, 4 dotted aqua bags.
clear teal bags contain 5 shiny maroon bags, 1 clear green bag.
dim brown bags contain 4 faded lavender bags, 5 striped lime bags, 1 dark aqua bag, 1 dark fuchsia bag.
muted beige bags contain 1 dim aqua bag, 4 plaid plum bags, 3 light white bags, 4 muted cyan bags.
muted blue bags contain 2 bright blue bags.
vibrant turquoise bags contain 3 muted crimson bags.
mirrored indigo bags contain 2 wavy lime bags, 5 bright olive bags, 5 bright black bags, 5 vibrant violet bags.
posh tomato bags contain 4 muted orange bags, 3 plaid white bags, 3 shiny tomato bags, 3 light beige bags.
bright gray bags contain 5 pale aqua bags, 3 shiny gold bags, 1 clear olive bag, 1 dull fuchsia bag.
pale green bags contain 4 light black bags, 3 posh purple bags, 2 clear chartreuse bags, 2 drab lime bags.
light orange bags contain 5 shiny chartreuse bags, 2 wavy blue bags, 2 wavy yellow bags.
light green bags contain 5 dark bronze bags, 4 light tan bags, 4 dim chartreuse bags.
shiny indigo bags contain 3 faded cyan bags.
dotted orange bags contain 2 wavy crimson bags, 3 dull green bags, 5 dark indigo bags.
dotted black bags contain 2 vibrant white bags.
plaid gold bags contain 3 mirrored bronze bags, 5 striped tan bags.
muted salmon bags contain 5 dull maroon bags, 1 vibrant tan bag, 1 dim purple bag, 4 dull chartreuse bags.
plaid salmon bags contain 5 dotted purple bags, 5 dim orange bags.
bright crimson bags contain 3 plaid maroon bags, 2 dim aqua bags, 5 dull magenta bags, 5 pale tomato bags.
dotted fuchsia bags contain 1 dark cyan bag, 1 striped magenta bag, 3 clear coral bags, 4 light purple bags.
dull blue bags contain 5 dim magenta bags, 1 mirrored maroon bag, 5 dark indigo bags.
dull orange bags contain 3 drab blue bags, 1 shiny beige bag.
muted red bags contain 5 clear brown bags, 5 striped turquoise bags, 3 dull fuchsia bags.
shiny coral bags contain 5 dark olive bags, 5 light blue bags.
drab violet bags contain 5 clear chartreuse bags, 2 posh orange bags, 3 pale purple bags.
clear blue bags contain 4 dull fuchsia bags, 4 faded purple bags, 3 mirrored plum bags.
vibrant beige bags contain 2 posh silver bags.
bright bronze bags contain 1 shiny yellow bag, 5 muted green bags, 3 dark gray bags.
bright orange bags contain no other bags.
light teal bags contain 3 mirrored magenta bags, 5 faded gray bags.
dark green bags contain 4 drab white bags, 2 drab green bags, 5 dotted coral bags, 1 mirrored black bag.
plaid silver bags contain 2 dotted bronze bags.
drab turquoise bags contain 3 bright white bags, 3 drab maroon bags.
dim magenta bags contain 5 dark fuchsia bags, 2 drab teal bags, 2 drab crimson bags, 2 dull fuchsia bags.
pale coral bags contain no other bags.
dull indigo bags contain 2 bright black bags, 1 drab lime bag, 5 light magenta bags, 1 faded orange bag.
dim indigo bags contain 5 dark maroon bags.
shiny gold bags contain 5 light black bags, 3 mirrored yellow bags, 5 muted plum bags.
faded lime bags contain 5 dark crimson bags, 3 shiny orange bags, 5 plaid tomato bags, 4 mirrored cyan bags.
faded cyan bags contain 3 shiny gold bags.
striped olive bags contain 4 vibrant red bags.
wavy white bags contain 5 posh silver bags, 5 mirrored gold bags, 5 pale black bags.
dim plum bags contain 5 dotted plum bags, 2 clear silver bags, 2 wavy bronze bags.
drab cyan bags contain 2 muted plum bags, 2 dotted bronze bags, 3 posh violet bags.
clear bronze bags contain 1 clear plum bag, 5 striped plum bags.
posh gray bags contain 5 clear chartreuse bags.
striped brown bags contain 3 dim olive bags, 1 light black bag, 4 vibrant crimson bags, 2 striped fuchsia bags.
wavy tomato bags contain 2 mirrored chartreuse bags.
dull white bags contain 1 mirrored brown bag, 2 dull green bags.
vibrant green bags contain 4 drab fuchsia bags.
dark tomato bags contain 3 light gray bags, 2 dull cyan bags, 4 striped silver bags, 5 dark fuchsia bags.
dim tomato bags contain 2 dark turquoise bags, 1 mirrored black bag, 3 posh maroon bags.
bright gold bags contain 5 pale aqua bags, 3 clear tomato bags, 1 dark yellow bag, 1 drab green bag.
dotted indigo bags contain 5 light purple bags, 2 plaid coral bags, 5 pale green bags.
plaid magenta bags contain 1 dotted bronze bag, 3 drab turquoise bags, 1 dark lime bag.
drab gray bags contain 3 dark gray bags, 5 clear turquoise bags.
dim gold bags contain 5 dark maroon bags.
dark black bags contain 5 dark violet bags, 2 dotted cyan bags.
posh chartreuse bags contain 2 posh magenta bags, 2 striped lime bags.
wavy beige bags contain 3 drab olive bags, 2 shiny beige bags, 1 faded purple bag.
plaid green bags contain 5 pale olive bags, 5 posh bronze bags, 4 bright tomato bags.
dotted blue bags contain 3 posh silver bags.
shiny cyan bags contain 1 mirrored turquoise bag, 1 striped beige bag, 5 bright silver bags, 3 light olive bags.
dark aqua bags contain 1 plaid black bag, 3 posh coral bags, 4 striped magenta bags, 2 mirrored turquoise bags.
drab salmon bags contain 2 striped bronze bags.
posh brown bags contain 3 bright white bags.
mirrored chartreuse bags contain 5 posh lime bags, 4 bright blue bags, 4 clear brown bags, 3 bright orange bags.
clear red bags contain 1 plaid beige bag, 2 posh brown bags, 2 shiny aqua bags.
wavy olive bags contain 4 clear maroon bags, 1 striped silver bag.
faded green bags contain 4 plaid red bags, 3 dim olive bags.
clear purple bags contain 1 plaid olive bag, 3 light chartreuse bags.
dull lime bags contain 4 dark tan bags, 1 light chartreuse bag, 5 vibrant silver bags.
dark red bags contain 5 faded orange bags.
wavy plum bags contain 5 dull teal bags, 3 clear maroon bags, 3 shiny tan bags.
dark white bags contain 3 muted red bags.
light turquoise bags contain 3 light black bags.
pale fuchsia bags contain 3 dim brown bags, 5 clear purple bags.
light maroon bags contain 3 mirrored turquoise bags.
striped fuchsia bags contain 1 dotted aqua bag.
mirrored magenta bags contain 5 striped white bags, 4 striped violet bags, 4 dull maroon bags, 5 striped indigo bags.
dark violet bags contain 2 mirrored black bags, 5 dotted fuchsia bags, 3 muted fuchsia bags.
muted cyan bags contain 3 muted salmon bags, 4 drab black bags, 2 posh green bags.
muted tan bags contain 4 posh coral bags, 2 bright fuchsia bags.
faded aqua bags contain 5 striped magenta bags, 1 dim aqua bag.
bright tomato bags contain 2 muted green bags, 1 light olive bag.
clear silver bags contain 4 wavy cyan bags, 3 bright orange bags, 5 mirrored coral bags, 3 light olive bags.
vibrant chartreuse bags contain 4 light gray bags.
dull olive bags contain 1 mirrored lavender bag, 4 dotted coral bags, 4 pale chartreuse bags, 1 dull coral bag.
pale gray bags contain 3 plaid orange bags.
wavy lime bags contain 1 bright orange bag, 2 wavy yellow bags, 2 light purple bags, 4 wavy indigo bags.
faded white bags contain 1 dotted violet bag, 1 dark maroon bag, 3 posh coral bags.
muted white bags contain 3 faded magenta bags.
wavy brown bags contain 4 vibrant yellow bags, 4 dull lavender bags.
clear violet bags contain 5 shiny tan bags.
clear olive bags contain 3 bright fuchsia bags, 5 dark maroon bags, 4 mirrored white bags, 5 shiny beige bags.
light beige bags contain 4 mirrored gray bags, 2 wavy brown bags, 3 pale blue bags, 4 striped silver bags.
plaid turquoise bags contain 1 vibrant aqua bag, 4 bright fuchsia bags.
posh coral bags contain 3 dark tomato bags.
wavy red bags contain 5 posh gray bags, 3 dim lime bags, 2 light tan bags, 3 bright blue bags.
wavy fuchsia bags contain 5 vibrant aqua bags.
faded turquoise bags contain 2 dark maroon bags, 1 pale indigo bag, 4 faded white bags.
clear indigo bags contain 4 pale purple bags, 5 dull green bags, 1 bright olive bag.
vibrant lime bags contain 1 light purple bag, 5 posh bronze bags, 5 drab blue bags, 1 bright black bag.
shiny bronze bags contain 1 plaid red bag.
vibrant plum bags contain 1 wavy black bag, 4 drab aqua bags, 5 dark cyan bags.
mirrored fuchsia bags contain 5 bright olive bags, 4 mirrored crimson bags, 1 dim salmon bag.
bright fuchsia bags contain 4 light gray bags.
bright silver bags contain 5 striped tan bags.
dotted gray bags contain 1 striped silver bag, 5 bright black bags, 2 mirrored yellow bags.
dark olive bags contain 5 striped lime bags, 1 bright black bag.
light lavender bags contain 1 dark bronze bag, 2 faded gold bags, 3 light orange bags.
mirrored green bags contain 2 faded orange bags.
faded blue bags contain 2 drab coral bags, 3 posh salmon bags.
vibrant red bags contain 3 bright cyan bags, 4 light aqua bags, 4 posh gray bags, 5 wavy purple bags.
dotted yellow bags contain 3 mirrored tan bags, 1 clear crimson bag, 3 light turquoise bags.
clear orange bags contain 4 pale coral bags, 3 posh silver bags, 2 dull fuchsia bags.
pale violet bags contain 5 light crimson bags.
mirrored beige bags contain 5 bright coral bags.
shiny tomato bags contain 4 dotted red bags, 2 plaid lavender bags, 5 dim orange bags.
muted bronze bags contain 3 striped tan bags, 3 faded orange bags, 2 faded maroon bags, 3 clear tomato bags.
muted fuchsia bags contain 1 dark maroon bag, 2 dotted bronze bags, 4 mirrored bronze bags, 1 faded cyan bag.
mirrored orange bags contain 2 plaid cyan bags, 5 wavy orange bags, 5 shiny aqua bags, 5 wavy tan bags.
light blue bags contain 1 mirrored chartreuse bag, 3 dim crimson bags.
bright turquoise bags contain 2 clear orange bags.
dark plum bags contain 3 wavy lime bags, 1 light tan bag, 3 light silver bags, 1 light lime bag.
wavy cyan bags contain 4 dull coral bags, 4 light olive bags.
striped cyan bags contain 4 dull lavender bags.
drab purple bags contain 4 shiny tomato bags, 4 bright orange bags, 4 mirrored gold bags.
faded indigo bags contain 5 mirrored indigo bags, 2 muted silver bags, 5 faded lime bags, 4 dim salmon bags.
faded orange bags contain 5 plaid chartreuse bags, 4 bright black bags, 5 light magenta bags, 4 wavy bronze bags.
wavy gold bags contain 1 shiny orange bag, 3 clear salmon bags, 3 plaid orange bags, 4 vibrant tan bags.
wavy blue bags contain 3 clear brown bags, 1 faded tomato bag, 5 drab green bags.
plaid violet bags contain 1 light blue bag, 5 drab purple bags.
wavy tan bags contain 4 dotted blue bags.
drab plum bags contain 2 muted silver bags, 5 shiny maroon bags.
drab fuchsia bags contain 2 muted maroon bags, 2 mirrored turquoise bags, 5 clear green bags, 3 light olive bags.
light violet bags contain 4 clear turquoise bags, 4 mirrored gold bags, 2 wavy chartreuse bags, 2 mirrored tan bags.
shiny maroon bags contain 1 plaid salmon bag, 4 pale brown bags, 1 dim orange bag, 1 wavy tomato bag.
drab crimson bags contain 2 dim gray bags, 5 dull fuchsia bags.
faded plum bags contain 2 striped turquoise bags, 5 light gray bags.
clear beige bags contain 4 faded orange bags, 2 mirrored black bags, 1 shiny red bag, 1 dark teal bag.
faded black bags contain 5 mirrored plum bags, 5 muted plum bags.
pale olive bags contain 4 muted yellow bags, 5 mirrored maroon bags.
plaid black bags contain 3 dark gray bags.
plaid fuchsia bags contain 1 wavy beige bag.
shiny white bags contain 3 posh orange bags, 5 posh blue bags, 4 faded white bags, 1 wavy crimson bag.
shiny lime bags contain 5 posh black bags, 2 mirrored bronze bags, 5 muted bronze bags, 2 posh violet bags.
shiny magenta bags contain 5 dark lime bags.
plaid tomato bags contain 5 wavy cyan bags, 3 clear brown bags, 3 dark olive bags, 4 vibrant white bags.
dark blue bags contain 3 mirrored black bags, 3 pale indigo bags, 3 dim cyan bags, 3 light olive bags.
muted lavender bags contain 1 dotted purple bag, 1 drab blue bag, 5 mirrored bronze bags, 3 striped violet bags.
faded silver bags contain 2 dim orange bags, 4 shiny chartreuse bags, 2 drab blue bags, 1 wavy violet bag.
shiny beige bags contain 4 muted green bags, 5 striped aqua bags, 2 dim black bags, 3 dull fuchsia bags.
vibrant indigo bags contain 3 pale red bags, 3 clear lime bags, 4 vibrant cyan bags, 2 pale tomato bags.
bright red bags contain 2 mirrored lime bags, 1 dim indigo bag, 5 bright black bags, 2 drab crimson bags.
light gray bags contain no other bags.
dim orange bags contain 4 clear chartreuse bags, 4 striped tan bags.
dull plum bags contain 4 faded aqua bags, 3 pale salmon bags, 1 posh gray bag.
dull crimson bags contain 5 posh red bags, 4 mirrored plum bags, 1 dull fuchsia bag.
pale blue bags contain 1 dotted red bag, 5 muted chartreuse bags, 3 clear green bags, 1 wavy beige bag.
dotted cyan bags contain 4 faded red bags, 2 bright gold bags.
mirrored lavender bags contain 1 vibrant white bag, 1 mirrored plum bag, 5 dotted black bags, 5 bright orange bags.
mirrored tomato bags contain 2 shiny chartreuse bags, 2 shiny bronze bags, 4 bright turquoise bags.
bright indigo bags contain 3 striped orange bags, 1 dotted lime bag, 1 shiny magenta bag, 2 light fuchsia bags.
drab brown bags contain 5 plaid magenta bags, 5 dim aqua bags, 4 vibrant aqua bags.
posh turquoise bags contain 3 bright tomato bags, 4 striped tomato bags, 5 dim turquoise bags.
shiny violet bags contain 1 drab gold bag, 5 plaid silver bags, 3 vibrant magenta bags.
bright green bags contain 5 dull aqua bags, 2 pale tomato bags, 1 posh lavender bag, 1 dim tomato bag.
striped tomato bags contain 4 bright salmon bags.
shiny black bags contain 3 drab aqua bags, 4 drab salmon bags, 1 dim turquoise bag.
dotted silver bags contain 4 plaid orange bags, 3 mirrored tan bags.
shiny red bags contain 4 dim lime bags, 3 posh bronze bags, 3 striped tomato bags, 2 vibrant aqua bags.
dim green bags contain 3 dotted blue bags, 4 faded cyan bags, 4 drab silver bags, 5 clear blue bags.
dull teal bags contain 4 striped orange bags, 5 bright coral bags, 4 bright gold bags.
posh beige bags contain 4 dark indigo bags.
clear green bags contain 3 drab blue bags, 2 dark maroon bags.
faded teal bags contain 4 mirrored maroon bags, 3 clear cyan bags, 4 plaid silver bags.
plaid maroon bags contain 2 plaid brown bags.
light yellow bags contain 3 dotted chartreuse bags.
dotted green bags contain 1 clear olive bag, 2 bright blue bags, 3 striped indigo bags, 3 dull indigo bags.
vibrant lavender bags contain 4 clear tomato bags, 1 posh tomato bag, 4 drab bronze bags.
dull red bags contain 1 dark red bag, 4 bright black bags.
dark chartreuse bags contain 2 clear turquoise bags, 2 clear coral bags, 2 vibrant magenta bags.
clear aqua bags contain 3 mirrored lime bags.
posh white bags contain 4 mirrored chartreuse bags, 1 light purple bag, 3 muted maroon bags, 2 pale olive bags.
pale cyan bags contain 3 plaid lime bags, 1 drab salmon bag.
mirrored yellow bags contain 4 light olive bags.
faded violet bags contain 2 muted red bags, 1 striped coral bag, 1 dark chartreuse bag, 3 vibrant aqua bags.
bright lime bags contain 1 muted chartreuse bag.
dotted teal bags contain 3 dark orange bags.
plaid beige bags contain 4 drab chartreuse bags, 5 clear orange bags, 1 dim orange bag, 4 dotted bronze bags.
muted violet bags contain 5 striped crimson bags, 3 dark gold bags, 4 muted magenta bags, 5 vibrant olive bags.
dotted turquoise bags contain 1 drab olive bag, 1 plaid turquoise bag.
dim yellow bags contain 4 dotted blue bags, 4 wavy teal bags.
light purple bags contain 1 mirrored yellow bag.
wavy turquoise bags contain 4 muted gold bags, 3 wavy orange bags, 3 clear tomato bags, 1 light tan bag.
vibrant silver bags contain 1 plaid red bag, 2 clear turquoise bags.
faded brown bags contain 5 faded gray bags, 3 drab maroon bags, 5 striped aqua bags.
posh indigo bags contain 1 wavy green bag, 5 dotted blue bags.
drab lavender bags contain 3 vibrant indigo bags, 2 faded black bags, 4 dull coral bags, 2 wavy lime bags.
vibrant blue bags contain 5 posh purple bags, 4 dark gold bags, 2 mirrored lavender bags.
posh blue bags contain 4 wavy bronze bags, 5 dull chartreuse bags, 1 muted teal bag, 3 bright black bags.
posh aqua bags contain 4 pale gold bags, 2 faded white bags.
bright yellow bags contain 1 wavy coral bag, 2 drab turquoise bags.
wavy lavender bags contain 2 faded red bags, 4 faded cyan bags.
dotted aqua bags contain 1 muted cyan bag, 2 muted black bags, 3 wavy chartreuse bags, 1 shiny magenta bag.
vibrant white bags contain 3 muted green bags, 2 bright tomato bags.
posh yellow bags contain 3 mirrored lime bags, 5 dark fuchsia bags, 1 posh red bag, 5 plaid cyan bags.
clear gray bags contain 4 wavy magenta bags, 3 shiny orange bags.
clear fuchsia bags contain 2 dull beige bags, 5 striped turquoise bags, 2 posh silver bags.
shiny tan bags contain 1 drab orange bag, 4 faded cyan bags, 5 dark teal bags.
faded bronze bags contain 2 shiny aqua bags.
pale indigo bags contain 1 plaid turquoise bag.
faded tomato bags contain 3 pale brown bags.
plaid indigo bags contain 1 dull beige bag.
dark salmon bags contain 3 wavy purple bags, 3 dull indigo bags, 4 dim blue bags, 3 dull green bags.
bright tan bags contain 2 posh bronze bags.
vibrant fuchsia bags contain 4 striped olive bags, 5 clear yellow bags, 5 muted fuchsia bags, 3 shiny plum bags.
dark gray bags contain 3 light gray bags.
posh magenta bags contain 5 bright blue bags.
mirrored bronze bags contain 5 bright olive bags, 4 light magenta bags.
posh cyan bags contain 4 light indigo bags, 2 dark aqua bags, 5 mirrored lime bags, 2 faded magenta bags.
light chartreuse bags contain 4 dark indigo bags, 2 wavy magenta bags, 5 dim white bags, 1 plaid bronze bag.
bright plum bags contain 4 dotted lime bags, 2 mirrored red bags, 1 plaid plum bag, 1 mirrored gold bag.
dark beige bags contain 4 mirrored white bags, 2 muted plum bags, 5 mirrored lime bags, 2 plaid teal bags.
light bronze bags contain 1 muted black bag.
bright magenta bags contain 3 striped lime bags.
dull green bags contain 2 light purple bags, 1 dull maroon bag, 2 dotted violet bags, 4 clear blue bags.
drab maroon bags contain 3 bright tomato bags.
muted gray bags contain 1 shiny plum bag, 2 posh fuchsia bags, 1 plaid black bag, 2 dim black bags.
pale aqua bags contain 5 mirrored plum bags, 1 dark fuchsia bag, 3 faded tomato bags, 1 striped aqua bag.
plaid yellow bags contain 2 plaid plum bags.
vibrant tan bags contain 3 light gray bags, 5 bright salmon bags, 3 pale green bags, 5 posh gray bags.
faded lavender bags contain 2 shiny yellow bags, 3 dotted black bags, 4 dotted purple bags.
muted aqua bags contain 3 mirrored blue bags, 2 plaid salmon bags.
wavy silver bags contain 3 posh lime bags, 2 striped tan bags.
pale lime bags contain 5 striped orange bags, 4 plaid turquoise bags, 1 dark lime bag, 5 muted cyan bags.
dark bronze bags contain 3 clear silver bags, 5 faded tomato bags, 5 light olive bags, 4 bright fuchsia bags.
dull gray bags contain 1 clear silver bag, 1 light purple bag.
clear tomato bags contain 3 wavy coral bags, 2 dim orange bags, 2 dim magenta bags.
dull chartreuse bags contain 2 plaid teal bags, 4 dotted purple bags, 1 faded tomato bag.
vibrant gold bags contain 2 striped aqua bags, 5 vibrant cyan bags, 2 dotted olive bags, 2 clear olive bags.
wavy green bags contain 5 dim lavender bags.
posh olive bags contain 3 striped blue bags, 4 striped beige bags, 4 dim violet bags, 4 muted blue bags.
vibrant brown bags contain 5 light purple bags, 1 bright orange bag.
faded yellow bags contain 2 dark salmon bags.
vibrant teal bags contain 5 vibrant brown bags, 5 shiny indigo bags.
drab black bags contain 4 bright magenta bags, 1 shiny green bag.
mirrored black bags contain 3 posh silver bags.
muted yellow bags contain 1 clear orange bag, 2 shiny gold bags, 4 wavy purple bags.
posh silver bags contain no other bags.
plaid coral bags contain 2 mirrored lavender bags, 5 drab lime bags, 4 pale brown bags, 4 dark maroon bags.
muted plum bags contain no other bags.
pale silver bags contain 2 dim brown bags, 1 light aqua bag, 4 shiny lavender bags.
dotted coral bags contain 2 dotted bronze bags, 1 clear violet bag, 1 vibrant magenta bag.
drab indigo bags contain 1 dotted crimson bag.
dim salmon bags contain 2 clear green bags, 4 muted chartreuse bags.
vibrant cyan bags contain 4 posh magenta bags, 4 clear violet bags.
muted olive bags contain 5 bright salmon bags, 2 dark silver bags.
drab blue bags contain 2 shiny yellow bags, 5 clear olive bags.
dark brown bags contain 3 muted cyan bags, 5 posh fuchsia bags.
dotted red bags contain 3 posh gray bags, 5 clear maroon bags, 3 posh fuchsia bags, 1 dark white bag.
light crimson bags contain 2 dark chartreuse bags.
wavy chartreuse bags contain 1 muted red bag, 5 dull chartreuse bags, 2 wavy bronze bags, 1 posh bronze bag.
plaid lime bags contain 3 pale white bags, 2 dull chartreuse bags, 3 plaid olive bags, 1 vibrant cyan bag.
striped gray bags contain 5 mirrored blue bags, 3 dark turquoise bags, 2 clear aqua bags, 5 drab cyan bags.
dull cyan bags contain no other bags.
dotted beige bags contain 5 dull cyan bags, 2 dull purple bags, 4 mirrored white bags, 3 vibrant olive bags.
pale turquoise bags contain 5 pale beige bags, 2 pale olive bags, 2 wavy coral bags, 5 light fuchsia bags.
muted tomato bags contain 1 vibrant olive bag, 1 bright purple bag, 3 pale turquoise bags, 3 striped coral bags.
pale white bags contain 4 dull gold bags, 5 wavy olive bags, 4 faded red bags, 2 plaid teal bags.
clear magenta bags contain 4 dark violet bags, 5 plaid chartreuse bags, 3 vibrant yellow bags.
vibrant tomato bags contain 2 posh lime bags, 4 drab orange bags, 1 striped turquoise bag.
striped green bags contain 1 dim purple bag, 3 dotted bronze bags, 4 bright bronze bags.
wavy salmon bags contain 1 shiny lime bag.
plaid purple bags contain 1 muted tomato bag, 2 shiny lavender bags, 5 light olive bags.
clear tan bags contain 2 striped plum bags, 1 striped fuchsia bag.
dull coral bags contain 5 bright orange bags, 5 faded purple bags, 5 plaid chartreuse bags, 3 muted green bags.
light tomato bags contain 4 faded tomato bags, 1 clear chartreuse bag, 2 plaid black bags, 2 posh plum bags.
dotted lime bags contain 3 bright lavender bags.
plaid blue bags contain 4 plaid black bags.
dull silver bags contain 1 wavy magenta bag, 2 mirrored fuchsia bags, 4 striped salmon bags.
dotted maroon bags contain 3 dull cyan bags, 5 plaid lavender bags, 3 bright gray bags.
bright black bags contain 4 mirrored plum bags, 2 drab blue bags, 3 light gray bags, 1 posh coral bag.
pale salmon bags contain 5 bright gray bags.
muted chartreuse bags contain 1 mirrored lavender bag.
pale lavender bags contain 4 dim black bags.
dull fuchsia bags contain 1 bright olive bag, 3 dull cyan bags, 3 bright tomato bags.
posh fuchsia bags contain 4 striped cyan bags, 1 shiny purple bag, 5 muted lavender bags.
dull tan bags contain 2 light magenta bags.
mirrored olive bags contain 5 clear maroon bags, 3 bright cyan bags, 2 vibrant plum bags.
plaid chartreuse bags contain 2 pale coral bags, 1 posh lime bag, 5 light olive bags, 2 bright orange bags.
muted teal bags contain 3 plaid teal bags.
dim violet bags contain 3 striped tomato bags, 1 dotted fuchsia bag.
striped yellow bags contain 2 mirrored brown bags, 3 faded cyan bags, 1 clear silver bag, 5 wavy orange bags.
faded salmon bags contain 5 striped coral bags.
striped turquoise bags contain no other bags.
dim white bags contain 5 clear coral bags.
dull violet bags contain 4 striped violet bags, 5 dotted olive bags, 4 pale gold bags, 2 vibrant olive bags.
posh plum bags contain 2 bright orange bags, 5 faded tomato bags, 3 pale brown bags, 1 posh silver bag.
wavy orange bags contain 3 dull maroon bags, 1 drab orange bag, 4 posh plum bags.
dotted purple bags contain 2 mirrored white bags.
dark indigo bags contain 3 muted green bags, 5 dark white bags, 4 drab olive bags, 5 vibrant tomato bags.
shiny silver bags contain 2 pale green bags.
shiny crimson bags contain 3 wavy chartreuse bags, 2 wavy olive bags.
dull salmon bags contain 5 plaid plum bags.
bright brown bags contain 4 clear tan bags.
wavy aqua bags contain 1 dotted tan bag, 4 bright turquoise bags, 1 wavy maroon bag, 4 shiny cyan bags.
mirrored cyan bags contain 1 pale green bag, 5 plaid chartreuse bags, 5 muted chartreuse bags, 1 faded purple bag.
dark yellow bags contain 5 mirrored crimson bags, 2 shiny beige bags, 5 mirrored brown bags, 4 muted aqua bags.
faded gray bags contain 5 wavy cyan bags, 2 dim olive bags, 5 wavy gray bags.
drab aqua bags contain 3 shiny purple bags, 2 dim gray bags, 3 wavy cyan bags.
vibrant salmon bags contain 2 light indigo bags, 4 pale maroon bags.
drab yellow bags contain 1 light purple bag, 5 muted fuchsia bags, 2 drab blue bags, 4 muted green bags.
dark maroon bags contain 4 posh coral bags.
drab orange bags contain 4 bright tomato bags, 4 faded purple bags, 5 pale brown bags.
dim teal bags contain 4 shiny gray bags.
dotted bronze bags contain 2 drab blue bags, 1 light magenta bag.
faded maroon bags contain 5 wavy cyan bags, 1 pale gold bag.
vibrant gray bags contain 3 dull coral bags, 4 faded lime bags, 3 mirrored turquoise bags.
wavy yellow bags contain 3 striped bronze bags.
mirrored white bags contain no other bags.
pale chartreuse bags contain 5 drab blue bags, 3 bright black bags, 1 mirrored lavender bag, 4 dotted magenta bags.
posh crimson bags contain 1 mirrored lavender bag, 1 clear cyan bag.
dim coral bags contain 2 posh brown bags.
striped crimson bags contain 3 dim gray bags, 1 light turquoise bag, 3 wavy bronze bags, 4 faded orange bags.
posh violet bags contain 1 dark teal bag, 4 posh red bags, 3 vibrant lime bags.
light salmon bags contain 3 plaid salmon bags.
plaid red bags contain 2 dull lavender bags, 1 posh plum bag, 4 faded cyan bags, 1 plaid turquoise bag.
pale maroon bags contain 1 plaid turquoise bag, 4 faded maroon bags, 4 shiny yellow bags, 1 pale purple bag.
striped magenta bags contain 2 posh turquoise bags, 5 wavy indigo bags, 4 plaid tomato bags, 3 dim lavender bags.
striped silver bags contain no other bags.
striped gold bags contain 5 bright brown bags, 1 dotted crimson bag, 2 bright olive bags.
clear turquoise bags contain 3 wavy teal bags, 2 muted red bags.
mirrored tan bags contain 5 dark yellow bags, 3 posh coral bags.
shiny blue bags contain 2 dull olive bags, 2 muted brown bags.
clear black bags contain 2 wavy teal bags, 5 plaid chartreuse bags, 4 dull coral bags, 5 dark yellow bags.
faded gold bags contain 5 muted teal bags, 3 bright white bags, 4 striped tan bags.
drab tan bags contain 1 faded black bag, 2 clear olive bags.
dark cyan bags contain 1 plaid black bag, 1 muted aqua bag, 5 bright fuchsia bags.
muted silver bags contain 1 clear olive bag, 5 striped indigo bags.
dim beige bags contain 1 muted tomato bag, 5 clear fuchsia bags, 1 faded coral bag.
striped red bags contain 3 plaid brown bags, 4 posh black bags, 2 dotted gray bags.
striped purple bags contain 5 light lavender bags, 2 dotted brown bags, 1 dull olive bag, 2 shiny aqua bags.
shiny orange bags contain 3 mirrored brown bags, 1 wavy bronze bag, 5 vibrant aqua bags.
striped salmon bags contain 4 bright silver bags.
shiny olive bags contain 4 pale gold bags, 5 drab indigo bags, 3 mirrored salmon bags, 2 muted gray bags.
pale bronze bags contain 3 clear cyan bags, 3 drab blue bags, 5 drab bronze bags, 4 shiny gray bags.
plaid aqua bags contain 5 pale yellow bags, 4 pale black bags, 3 muted red bags.
faded coral bags contain 3 dark bronze bags, 5 striped silver bags, 5 clear olive bags, 2 wavy gray bags.
striped bronze bags contain 4 posh orange bags.
bright violet bags contain 5 light white bags, 1 dull olive bag, 5 drab fuchsia bags, 3 dim chartreuse bags.
wavy bronze bags contain 2 dark tomato bags, 2 muted red bags, 1 drab orange bag.
pale beige bags contain 5 muted lavender bags, 1 vibrant aqua bag, 4 drab lime bags.
dim crimson bags contain 4 plaid plum bags.
light silver bags contain 3 shiny yellow bags, 4 dull fuchsia bags, 4 dark chartreuse bags, 1 bright orange bag.
dark coral bags contain 1 clear indigo bag, 1 muted gold bag, 5 pale lime bags.
striped plum bags contain 5 plaid white bags, 3 pale gold bags, 3 pale yellow bags, 2 dim orange bags.
light tan bags contain 5 dark crimson bags, 1 clear silver bag, 2 striped tomato bags, 1 vibrant magenta bag.
shiny brown bags contain 4 bright bronze bags.
faded red bags contain 2 dotted bronze bags.
dim gray bags contain 1 pale gold bag, 5 shiny orange bags.
mirrored plum bags contain 2 muted plum bags, 1 posh silver bag.
shiny yellow bags contain 4 faded black bags, 4 light olive bags.
posh bronze bags contain 3 posh orange bags.
clear lime bags contain 5 mirrored lavender bags, 1 dark tomato bag, 4 dim aqua bags, 1 pale purple bag.
drab red bags contain 3 dark tan bags, 2 shiny maroon bags, 2 mirrored purple bags, 5 dotted orange bags.
dull purple bags contain 1 striped yellow bag, 3 faded cyan bags, 5 pale red bags, 4 plaid green bags.
muted magenta bags contain 1 drab yellow bag, 1 dark lavender bag.
dim fuchsia bags contain 4 pale aqua bags, 3 mirrored indigo bags, 2 wavy lime bags.
muted black bags contain 3 bright turquoise bags, 3 plaid cyan bags, 5 dim cyan bags.
dark lime bags contain 1 posh orange bag.
drab bronze bags contain 1 pale tomato bag, 4 light purple bags, 1 light olive bag, 4 posh silver bags.
posh tan bags contain 5 bright gold bags.
dim lavender bags contain 1 mirrored white bag, 4 posh lime bags, 3 dark fuchsia bags.
dark fuchsia bags contain no other bags.
muted maroon bags contain 2 bright white bags, 4 dark salmon bags, 4 posh gray bags, 4 posh plum bags.
dotted crimson bags contain 2 plaid salmon bags.
drab olive bags contain 2 dull lavender bags, 3 dark tomato bags.
wavy crimson bags contain 3 clear orange bags, 5 dull maroon bags.
dark silver bags contain 1 muted red bag, 5 dim bronze bags.
dull magenta bags contain 2 mirrored cyan bags, 2 mirrored plum bags, 1 drab olive bag.
pale magenta bags contain 4 dark olive bags, 1 wavy teal bag.
plaid crimson bags contain 1 dim purple bag, 1 shiny gold bag, 5 shiny tan bags, 2 striped silver bags.
vibrant magenta bags contain 3 striped turquoise bags.
light brown bags contain 2 clear magenta bags, 2 light lime bags.
drab lime bags contain 2 striped aqua bags.
light aqua bags contain 2 dim orange bags, 5 mirrored brown bags, 4 vibrant tomato bags.
dotted white bags contain 3 faded plum bags, 1 striped lime bag.
muted orange bags contain 4 clear purple bags, 5 light indigo bags, 1 plaid bronze bag.
dark turquoise bags contain 2 clear cyan bags.
striped beige bags contain 3 wavy yellow bags, 2 clear brown bags, 1 faded plum bag, 2 dotted bronze bags.
bright chartreuse bags contain 2 dim cyan bags, 2 faded lavender bags, 3 muted yellow bags, 1 dotted turquoise bag.
bright beige bags contain 3 vibrant silver bags, 3 faded bronze bags, 4 bright lime bags, 5 plaid lavender bags.
pale orange bags contain 1 striped tomato bag, 3 pale brown bags, 5 plaid bronze bags, 4 dark salmon bags.
dim black bags contain 4 striped turquoise bags, 2 plaid chartreuse bags, 5 posh red bags, 1 bright tomato bag.
plaid olive bags contain 5 dark indigo bags, 1 dark teal bag.
plaid orange bags contain 5 muted plum bags, 4 dark tomato bags, 5 dull crimson bags.
dark purple bags contain 3 pale beige bags, 4 pale gold bags, 1 vibrant blue bag.
mirrored salmon bags contain 3 plaid red bags, 3 dark plum bags.
vibrant aqua bags contain 3 clear maroon bags, 1 striped silver bag, 5 shiny gold bags, 3 faded tomato bags.
dull lavender bags contain 3 faded white bags, 1 dim lavender bag, 2 dull fuchsia bags.
posh green bags contain 5 dim orange bags.
clear salmon bags contain 4 bright black bags, 5 dotted plum bags, 2 striped tomato bags.
mirrored violet bags contain 2 wavy black bags, 5 dotted gold bags, 3 posh brown bags.
faded magenta bags contain 2 vibrant crimson bags, 5 drab orange bags, 1 dark gray bag, 4 striped coral bags.
light plum bags contain 4 muted plum bags.
dotted tan bags contain 1 wavy gray bag.
dim tan bags contain 5 mirrored bronze bags, 3 drab olive bags, 2 wavy olive bags, 3 dark tan bags.
striped indigo bags contain 2 drab green bags, 1 light olive bag, 5 bright orange bags.
vibrant orange bags contain 4 dull fuchsia bags, 1 shiny violet bag.
dim lime bags contain 5 striped silver bags.
plaid gray bags contain 4 dark tan bags, 3 dark magenta bags, 2 drab black bags, 3 faded bronze bags.
faded chartreuse bags contain 2 clear beige bags, 2 light beige bags.
striped orange bags contain 5 mirrored coral bags, 4 light gray bags, 2 mirrored white bags.
dim silver bags contain 3 posh coral bags, 4 dotted bronze bags, 4 muted yellow bags, 1 faded gray bag.
dotted gold bags contain 1 posh gray bag.
dotted olive bags contain 5 faded aqua bags, 4 faded brown bags, 2 dim salmon bags.
posh maroon bags contain 4 striped aqua bags, 1 wavy yellow bag, 1 mirrored crimson bag.
dark orange bags contain 3 mirrored fuchsia bags, 1 light magenta bag, 2 muted bronze bags, 2 wavy blue bags.
pale purple bags contain 4 dotted blue bags.
muted indigo bags contain 1 dark fuchsia bag, 2 posh fuchsia bags, 5 plaid turquoise bags, 2 mirrored cyan bags.
faded tan bags contain 1 drab fuchsia bag, 3 dark beige bags.
bright cyan bags contain 3 dotted blue bags, 5 muted plum bags, 2 vibrant crimson bags, 3 dark gold bags.
mirrored gray bags contain 3 pale lavender bags, 2 shiny blue bags, 4 dark salmon bags.
dotted violet bags contain 4 posh red bags, 4 bright tomato bags, 3 muted plum bags.
pale tan bags contain 2 clear crimson bags, 4 drab lime bags.
dull tomato bags contain 2 drab green bags, 5 dark crimson bags.
wavy coral bags contain 3 light aqua bags, 3 bright blue bags, 1 posh gray bag.
dull aqua bags contain 4 drab brown bags.
clear white bags contain 5 dark indigo bags, 5 dark maroon bags, 1 striped orange bag, 2 shiny gold bags.
dotted magenta bags contain 3 light olive bags, 4 dark indigo bags, 3 dotted blue bags, 3 striped lime bags.
plaid cyan bags contain 3 wavy olive bags, 5 dim black bags, 4 dotted bronze bags, 2 striped tomato bags.
muted lime bags contain 3 pale turquoise bags, 1 posh white bag.
wavy teal bags contain 2 dim orange bags.
muted gold bags contain 1 pale aqua bag, 4 shiny beige bags, 2 light olive bags.
wavy magenta bags contain 2 wavy purple bags, 5 wavy tomato bags.
dim maroon bags contain 5 plaid red bags.
drab green bags contain 5 light gray bags, 4 clear maroon bags, 2 dark indigo bags.
plaid lavender bags contain 5 drab blue bags.
light white bags contain 4 light tan bags.
wavy gray bags contain 1 posh coral bag.
mirrored teal bags contain 5 clear brown bags, 4 bright magenta bags, 1 drab brown bag, 2 dull gold bags.
light black bags contain 1 posh coral bag, 4 dotted black bags, 4 posh lime bags, 4 bright blue bags.
shiny gray bags contain 2 dotted blue bags, 5 striped turquoise bags, 4 pale aqua bags, 1 dim black bag.
posh lavender bags contain 1 wavy yellow bag, 2 dotted tan bags, 3 dull lavender bags.
faded beige bags contain 5 posh brown bags, 1 vibrant indigo bag, 2 light cyan bags, 1 clear aqua bag.
clear lavender bags contain 1 dull salmon bag.
mirrored coral bags contain 2 light olive bags, 5 clear olive bags, 2 pale tomato bags.
bright white bags contain 5 wavy yellow bags, 5 wavy bronze bags, 1 wavy olive bag, 5 muted red bags.
dark lavender bags contain 5 pale tomato bags, 4 faded white bags.
light red bags contain 3 posh turquoise bags, 3 dull indigo bags, 3 wavy silver bags, 2 drab salmon bags.
vibrant violet bags contain 1 shiny orange bag.
clear gold bags contain 1 drab magenta bag, 4 plaid tan bags, 2 vibrant yellow bags.
striped white bags contain 5 light salmon bags, 1 mirrored purple bag.
vibrant maroon bags contain 4 striped coral bags.
bright olive bags contain 2 light gray bags, 1 posh silver bag, 2 bright orange bags, 1 dark fuchsia bag.
mirrored gold bags contain 1 drab black bag.
shiny green bags contain 4 posh purple bags.
dim aqua bags contain 4 pale gold bags, 1 dull coral bag, 3 faded teal bags, 2 pale yellow bags.
dotted tomato bags contain 4 dull salmon bags.
faded olive bags contain 5 muted green bags, 1 drab crimson bag.
striped teal bags contain 4 dotted green bags, 5 muted aqua bags.
pale crimson bags contain 1 striped lime bag, 4 mirrored tan bags, 2 clear lavender bags.
shiny salmon bags contain 2 muted salmon bags, 5 shiny cyan bags, 4 faded red bags, 5 light tan bags.
plaid brown bags contain 5 pale orange bags.
posh gold bags contain 3 shiny crimson bags.
wavy violet bags contain 5 muted olive bags.
drab white bags contain 3 posh brown bags, 2 striped cyan bags, 1 clear coral bag.
light magenta bags contain 2 clear maroon bags, 3 light gray bags, 2 dotted black bags, 4 bright fuchsia bags.
bright aqua bags contain 4 plaid cyan bags, 2 clear black bags.
muted green bags contain 5 striped silver bags, 5 bright orange bags.
dim cyan bags contain 5 plaid brown bags, 3 striped tan bags.
vibrant olive bags contain 5 dark yellow bags.
dark tan bags contain 1 muted lavender bag, 2 mirrored turquoise bags, 1 dim lime bag, 1 dull olive bag.
mirrored brown bags contain 3 pale brown bags.
mirrored aqua bags contain 1 mirrored plum bag, 3 dark maroon bags.
clear cyan bags contain 1 dull cyan bag, 2 dark tomato bags, 4 pale brown bags.
dull yellow bags contain 4 wavy coral bags, 4 striped tan bags, 3 muted chartreuse bags.
mirrored maroon bags contain 3 plaid tomato bags, 4 shiny purple bags, 1 plaid lavender bag, 5 light gray bags.
drab coral bags contain 2 dark turquoise bags, 2 clear crimson bags, 4 drab lime bags, 5 dull crimson bags.
muted turquoise bags contain 5 posh purple bags, 2 dim magenta bags.
dull turquoise bags contain 4 striped aqua bags, 2 light fuchsia bags, 3 pale turquoise bags, 2 faded maroon bags.
bright maroon bags contain 1 muted red bag, 3 faded black bags.
wavy indigo bags contain 3 shiny beige bags, 5 dim lavender bags, 2 striped tan bags.
pale red bags contain 1 vibrant brown bag, 3 faded black bags, 4 posh turquoise bags, 5 plaid brown bags.
drab teal bags contain 1 faded purple bag, 2 mirrored blue bags, 2 vibrant tomato bags, 1 pale purple bag.
faded purple bags contain 1 vibrant white bag.
striped maroon bags contain 1 mirrored brown bag, 1 shiny black bag, 5 dotted lime bags.
dotted chartreuse bags contain 3 dim blue bags.
posh black bags contain 3 light aqua bags, 5 bright orange bags, 1 plaid plum bag, 5 plaid fuchsia bags.
pale gold bags contain 1 faded tomato bag, 2 dark tomato bags, 3 dotted blue bags.
muted brown bags contain 1 plaid brown bag.
vibrant purple bags contain 1 shiny gray bag, 5 dull green bags.
mirrored purple bags contain 1 bright gray bag, 2 plaid plum bags, 5 dotted chartreuse bags, 2 posh lime bags.
bright coral bags contain 3 dim blue bags.
dark magenta bags contain 4 muted maroon bags, 2 dark olive bags, 1 dull olive bag.
pale brown bags contain 2 faded purple bags, 1 muted green bag, 3 dark fuchsia bags.
drab gold bags contain 3 posh bronze bags, 2 plaid tomato bags.
pale teal bags contain 1 clear turquoise bag, 4 muted yellow bags, 1 posh tomato bag, 3 vibrant orange bags.
vibrant black bags contain 1 faded tomato bag, 5 dim white bags, 2 drab aqua bags, 1 vibrant silver bag.
dull black bags contain 3 pale yellow bags, 4 clear aqua bags, 1 shiny yellow bag, 3 faded maroon bags.
drab tomato bags contain 2 dim purple bags, 5 plaid fuchsia bags.
bright blue bags contain 5 posh silver bags, 4 dull cyan bags, 4 light olive bags, 1 mirrored white bag.
muted crimson bags contain 3 bright orange bags, 1 pale tomato bag, 3 posh yellow bags, 4 shiny purple bags.
shiny plum bags contain 1 dim turquoise bag.
clear maroon bags contain 1 posh red bag, 2 light olive bags, 1 dotted purple bag.
drab chartreuse bags contain 4 dotted green bags, 1 dark red bag.
dull brown bags contain 5 dim lavender bags, 4 bright tomato bags, 5 drab crimson bags, 1 vibrant tomato bag.
wavy purple bags contain 4 dim lavender bags, 2 plaid plum bags, 2 dim lime bags, 2 striped cyan bags.
vibrant coral bags contain 1 striped crimson bag, 3 drab cyan bags, 5 vibrant yellow bags, 2 dotted teal bags.
clear crimson bags contain 4 dotted purple bags, 5 faded cyan bags.
clear coral bags contain 4 clear chartreuse bags, 2 bright olive bags.
light cyan bags contain 3 shiny violet bags, 3 dotted blue bags, 3 drab beige bags.
light fuchsia bags contain 2 plaid lavender bags, 4 dull green bags, 2 plaid salmon bags.
dull bronze bags contain 3 bright magenta bags, 1 bright black bag, 2 wavy lime bags.
light gold bags contain 5 bright indigo bags.
drab magenta bags contain 5 faded fuchsia bags, 2 dim black bags, 5 dim crimson bags, 5 dotted fuchsia bags.
mirrored lime bags contain 2 mirrored white bags, 1 shiny gray bag.
dim olive bags contain 4 bright blue bags, 1 faded tomato bag.
striped tan bags contain 4 dark tomato bags, 4 mirrored coral bags, 2 mirrored lavender bags.
muted coral bags contain 4 pale salmon bags, 3 faded purple bags.
striped violet bags contain 4 dark gray bags, 3 posh coral bags, 2 striped turquoise bags, 4 bright fuchsia bags.
clear plum bags contain 3 bright cyan bags, 1 drab salmon bag, 2 vibrant brown bags.
clear brown bags contain 1 mirrored plum bag.
light indigo bags contain 4 drab blue bags, 5 mirrored chartreuse bags, 2 muted red bags, 2 dark beige bags.
dim purple bags contain 4 wavy teal bags.
clear yellow bags contain 4 clear maroon bags, 3 bright olive bags, 5 shiny gray bags, 2 bright orange bags.
pale tomato bags contain 2 dotted violet bags, 2 dark teal bags.
bright teal bags contain 2 shiny maroon bags, 2 muted indigo bags.
pale plum bags contain 5 posh gray bags, 3 shiny indigo bags, 3 wavy olive bags, 1 pale white bag.
dark teal bags contain 3 muted plum bags, 4 faded plum bags, 1 wavy bronze bag.
vibrant crimson bags contain 5 dark tomato bags, 2 dark white bags, 5 posh red bags.
shiny aqua bags contain 1 pale coral bag.
dim chartreuse bags contain 2 wavy teal bags, 5 mirrored black bags, 5 mirrored bronze bags, 4 muted lavender bags.
drab beige bags contain 2 faded coral bags, 5 muted tan bags, 5 plaid cyan bags.
bright purple bags contain 4 muted red bags, 5 wavy beige bags, 4 clear coral bags.
striped black bags contain 3 light red bags, 2 plaid chartreuse bags.
striped chartreuse bags contain 2 dotted gray bags, 2 wavy olive bags, 1 muted lavender bag.
dull gold bags contain 5 wavy olive bags, 2 posh plum bags, 4 shiny gold bags.
striped lavender bags contain 5 dim teal bags, 3 light blue bags.
striped lime bags contain 4 light black bags, 5 striped turquoise bags, 5 wavy cyan bags.`
	return strings.Split(input, "\n")
}
