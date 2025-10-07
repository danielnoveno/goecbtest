package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
	"log"
)

type Contact struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstname,omitempty"`
	LastName  string             `bson:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
}

type ContactCollection struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewContactCollection(contactcollection *mongo.Collection, ctx context.Context) *ContactCollection {
	return &ContactCollection{
		collection: contactcollection,
		ctx:        ctx,
	}
}

func (c *ContactCollection) CreateContact(contact Contact) error {
	_, err := c.collection.InsertOne(c.ctx, contact)
	return err
}

func (c *ContactCollection) GetAll() ([]Contact, error) {
	var contacts []Contact
	cursor, err := c.collection.Find(c.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.ctx, &contacts); err != nil {
		panic(err)
	}

	cursor.Close(c.ctx)

	return contacts, nil
}

func (c *ContactCollection) UpdateContact(contact Contact) error {
	filter := bson.M{"_id": contact.ID}
	_, err := c.collection.ReplaceOne(c.ctx, filter, contact)
	return err
}

func (c *ContactCollection) DeleteContact(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := c.collection.DeleteOne(c.ctx, filter)
	return err
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Contact Data")

	var loadedData []Contact

	ctx := context.TODO()

	r, _ := fyne.LoadResourceFromPath("E:/Magang - Noveno/goecbtest_3/goecbtest/assets/images/cat.png")
    myWindow.SetIcon(r)

	opts := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	log.Println("mongo connection established")

	contactcollection := client.Database("contactdb").Collection("contacts")
	cc := NewContactCollection(contactcollection, ctx)

	loadedData, err = cc.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	var list *widget.List

	list = widget.NewList(
		func() int { return len(loadedData) },
		func() fyne.CanvasObject {
			firstNameLabel := widget.NewLabel("First Name")
			lastNameLabel := widget.NewLabel("Last Name")
			emailLabel := widget.NewLabel("Email")
			phoneLabel := widget.NewLabel("Phone")
			contactContainer := container.NewGridWithColumns(4, firstNameLabel, lastNameLabel, emailLabel, phoneLabel)
			edtb := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			delb := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			buttonContainer := container.NewHBox(edtb, delb)
			return container.NewBorder(nil, nil, nil, buttonContainer, contactContainer)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)

			contactContainer := c.Objects[0].(*fyne.Container)
			buttonContainer := c.Objects[1].(*fyne.Container)

			firstNameLabel := contactContainer.Objects[0].(*widget.Label)
			lastNameLabel := contactContainer.Objects[1].(*widget.Label)
			emailLabel := contactContainer.Objects[2].(*widget.Label)
			phoneLabel := contactContainer.Objects[3].(*widget.Label)

			edtb := buttonContainer.Objects[0].(*widget.Button)
			delb := buttonContainer.Objects[1].(*widget.Button)

			contactID := loadedData[i].ID

			firstNameLabel.SetText(loadedData[i].FirstName)
			lastNameLabel.SetText(loadedData[i].LastName)
			emailLabel.SetText(loadedData[i].Email)
			phoneLabel.SetText(loadedData[i].Phone)

			edtb.OnTapped = func() {

				entryFirstName := widget.NewEntry()
				entryLastName := widget.NewEntry()
				entryEmail := widget.NewEntry()
				entryPhone := widget.NewEntry()

				firstNameForm := widget.NewFormItem("First Name", entryFirstName)
				lastNameForm := widget.NewFormItem("Last Name", entryLastName)
				emailForm := widget.NewFormItem("Email", entryEmail)
				phoneForm := widget.NewFormItem("Phone", entryPhone)

				formItems := []*widget.FormItem{firstNameForm, lastNameForm, emailForm, phoneForm}

				dialog1 := dialog.NewForm("Edit Contact", "Save", "Cancel", formItems, func(b bool) {
					if b {
						contact := Contact{
							ID:        contactID,
							FirstName: entryFirstName.Text,
							LastName:  entryLastName.Text,
							Email:     entryEmail.Text,
							Phone:     entryPhone.Text,
						}
						if err := cc.UpdateContact(contact); err != nil {
							log.Fatal(err)
						}
						loadedData, err = cc.GetAll()
						if err != nil {
							log.Fatal(err)
						}
					}
					list.Refresh()
				}, myWindow)

				entryFirstName.SetText(firstNameLabel.Text)
				entryLastName.SetText(lastNameLabel.Text)
				entryEmail.SetText(emailLabel.Text)
				entryPhone.SetText(phoneLabel.Text)

				dialog1.Resize(fyne.NewSize(500, 300))

				dialog1.Show()

			}

			delb.OnTapped = func() {

				dialog1 := dialog.NewConfirm(
					"Delete Contact",
					"Do you wish to delete this contact?",
					func(b bool) {
						if b {
							if err := cc.DeleteContact(contactID); err != nil {
								log.Fatal(err)
							}
							loadedData, err = cc.GetAll()
							if err != nil {
								log.Fatal(err)
							}
						}
						list.Refresh()
					},
					myWindow,
				)

				dialog1.Resize(fyne.NewSize(300, 200))
				dialog1.Show()

			}

		})

	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
	}

	add := widget.NewButton("Add Contact", func() {

		entryFirstName := widget.NewEntry()
		entryLastName := widget.NewEntry()
		entryEmail := widget.NewEntry()
		entryPhone := widget.NewEntry()

		firstNameForm := widget.NewFormItem("First Name", entryFirstName)
		lastNameForm := widget.NewFormItem("Last Name", entryLastName)
		emailForm := widget.NewFormItem("Email", entryEmail)
		phoneForm := widget.NewFormItem("Phone", entryPhone)

		formItems := []*widget.FormItem{firstNameForm, lastNameForm, emailForm, phoneForm}

		dialogAdd := dialog.NewForm("Add Contact", "Add", "Cancel", formItems, func(b bool) {
			if b {
				c := Contact{
					FirstName: entryFirstName.Text,
					LastName:  entryLastName.Text,
					Email:     entryEmail.Text,
					Phone:     entryPhone.Text,
				}

				if err := cc.CreateContact(c); err != nil {
					log.Fatal(err)
				}
				loadedData, err = cc.GetAll()
				if err != nil {
					log.Fatal(err)
				}
				list.Refresh()
			}
		}, myWindow)

		dialogAdd.Resize(fyne.NewSize(500, 300))

		dialogAdd.Show()

	})

	exit := widget.NewButton("Quit", func() {

		myWindow.Close()
	})

	myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}