package game

import (
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/kettek/morogue/client/embed"
	"github.com/kettek/morogue/client/ifs"
	"github.com/kettek/morogue/game"
)

func addObjectInfo(ctx ifs.RunContext, object game.Object, arch game.Archetype, container *widget.Container) {
	noneColor := color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	lightColor := color.NRGBA{R: 100, G: 100, B: 200, A: 255}
	mediumColor := color.NRGBA{R: 200, G: 200, B: 100, A: 255}
	heavyColor := color.NRGBA{R: 200, G: 100, B: 100, A: 255}

	unarmedColor := color.NRGBA{R: 250, G: 150, B: 50, A: 255}
	rangedColor := color.NRGBA{R: 50, G: 250, B: 50, A: 255}
	meleeColor := color.NRGBA{R: 250, G: 50, B: 50, A: 255}

	switch a := arch.(type) {
	case game.WeaponArchetype:
		weaponColor := noneColor
		if a.WeaponType == game.WeaponTypeRange {
			weaponColor = rangedColor
		} else if a.WeaponType == game.WeaponTypeMelee {
			weaponColor = meleeColor
		} else if a.WeaponType == game.WeaponTypeUnarmed {
			weaponColor = unarmedColor
		}

		title := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(fmt.Sprintf("%s", a.Title), ctx.UI.BodyCopyFace, color.White))
		values := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(fmt.Sprintf("%s %s", a.RangeString(), a.WeaponType), ctx.UI.BodyCopyFace, weaponColor))
		desc := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(a.Description, ctx.UI.BodyCopyFace, color.NRGBA{128, 128, 128, 255}))

		weaponLine := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Direction(widget.DirectionHorizontal))),
		)
		graphic := widget.NewGraphic(
			widget.GraphicOpts.Image(embed.IconOffense),
			widget.GraphicOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		)
		weaponLine.AddChild(values)
		weaponLine.AddChild(graphic)

		container.AddChild(title)
		container.AddChild(weaponLine)
		container.AddChild(desc)
	case game.ArmorArchetype:
		armorColor := noneColor
		if a.ArmorType == game.ArmorTypeLight {
			armorColor = lightColor
		} else if a.ArmorType == game.ArmorTypeMedium {
			armorColor = mediumColor
		} else if a.ArmorType == game.ArmorTypeHeavy {
			armorColor = heavyColor
		}

		title := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(fmt.Sprintf("%s", a.Title), ctx.UI.BodyCopyFace, color.White))
		values := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(fmt.Sprintf("%s %s", a.RangeString(), a.ArmorType), ctx.UI.BodyCopyFace, armorColor))
		desc := widget.NewText(widget.TextOpts.ProcessBBCode(true), widget.TextOpts.Text(a.Description, ctx.UI.BodyCopyFace, color.NRGBA{128, 128, 128, 255}))

		armorLine := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Direction(widget.DirectionHorizontal))),
		)
		graphic := widget.NewGraphic(
			widget.GraphicOpts.Image(embed.IconDefense),
			widget.GraphicOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		)
		armorLine.AddChild(values)
		armorLine.AddChild(graphic)

		container.AddChild(title)
		container.AddChild(armorLine)
		container.AddChild(desc)
	case game.ItemArchetype:
		// TODO
	}

}
