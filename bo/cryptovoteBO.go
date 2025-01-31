package bo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/dao"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/interfaces"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	In go, there is no explicit declaration (keyword) to set an inheritance relationship between classes,
	as occurs in Java or Python, for example. There is the keyword interface, but this only indicates that
	it is an abstract struct. A class that wants to be an implementation of an interface-type struct needs
	to implement all the methods of the given interface.

	bo.CryptoVoteBO implements all methods of interfaces.InterfaceCryptoVoteBO
*/
type CryptoVoteBO struct {
	ImplDAO interfaces.InterfaceCryptoVoteDAO
}

/*
	Used as default to creates an instance with an arrowed DAO implementation, to bank-related actions at
	Mongodb.
*/
func NewCryptoVoteBO(iDAO interfaces.InterfaceCryptoVoteDAO) CryptoVoteBO {
	// creates an instance da struct CryptoVoteBO
	var bo CryptoVoteBO

	// creates an instance of DAO to manage bank-related actions at Mongodb
	// this struct implements all methods of interfaces.InterfaceCryptoVoteDAO
	if !(iDAO != nil) {
		iDAO = dao.NewCryptoVoteDAO()
	}

	// set the concrete implementation to interfaces.InterfaceCryptoVoteDAO
	bo = CryptoVoteBO{
		ImplDAO: iDAO,
	}

	// return bo ready to perform bank-related actions
	return bo
}

/*
	Used at test, to enable a Mock of InterfaceCryptoVoteDAO, to simulate bank-related actions at Mongodb.
*/
func (bo CryptoVoteBO) SetCryptoVoteDAO(iDAO interfaces.InterfaceCryptoVoteDAO) {
	bo.ImplDAO = iDAO
}

/*
	CreateCryptoVote validates data input before save. Return a model.CryptoVote stored at database.
*/
func (bo CryptoVoteBO) CreateCryptoVote(cryptoVote model.CryptoVote) (model.CryptoVote, error) {
	var savedCryptoVote model.CryptoVote

	cryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(cryptoVote.Name)))
	cryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(cryptoVote.Symbol))
	cryptoVote.Sum = cryptoVote.Qtd_Upvote - cryptoVote.Qtd_Downvote
	cryptoVote.SumAbsolute = utils.Abs(cryptoVote.Sum)

	// uses function from package bo
	validate, err := bo.ValidateCryptoVote(cryptoVote)
	if !validate || err != nil {
		return cryptoVote, err
	} else {

		// uses function from package dao
		savedCryptoVote, err = bo.ImplDAO.Create(cryptoVote)
		if err != nil {
			z := "[cryptovote.mongodb] Problemas na execução de bo.cryptoVoteDAO.Create: " + err.Error()
			err = errors.New(z)
			return cryptoVote, err
		}
	}
	// returns the new CryptoCurrency
	return savedCryptoVote, err
}

/*
	RetrieveAllCryptoVoteByFilter using filter to find many itens. Returns a list of
	model.CryptoVote stored at database or nil, if filter mismatch.

	// all data from model
	var filterCryptoVote = model.FilterCryptoVote{
		Name: "",
		Symbol: "",
	}

	// all data Name = Klever
	var filterCryptoVote = model.FilterCryptoVote{
		Name: "Klever",
		Symbol: "",
	}

	// all data Symbol = KLV
	var filterCryptoVote = model.FilterCryptoVote{
		Name: "",
		Symbol: "KLV",
	}
*/
func (bo CryptoVoteBO) RetrieveAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) ([]model.CryptoVote, error) {
	var listCryptoVote []model.CryptoVote

	filter, err := utils.MarshalFilterCryptoVoteToBsonFilter(filterCryptoVote)
	if err != nil {
		return listCryptoVote, err
	}

	// uses function from package dao
	listCryptoVote, err = bo.ImplDAO.FindMany(filter)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using bo.cryptoVoteDAO.FindMany: " + err.Error()
		err = errors.New(z)
		return listCryptoVote, err
	}
	return listCryptoVote, err
}

/*
	RetrieveOneCryptoVote creates a filter using args (name or symbol) to find one item.
	At least on filter shouldn't be empty. Returns a model.CryptoVote stored at database or nil
	if filter mismatch.
*/
func (bo CryptoVoteBO) RetrieveOneCryptoVote(name string, symbol string) (model.CryptoVote, error) {
	return bo.retrieveOneCryptoVote(name, symbol)
}

func (bo CryptoVoteBO) retrieveOneCryptoVote(name string, symbol string) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote
	var err error

	// creates a filter using args
	var filterCryptoVote = model.FilterCryptoVote{
		Name:   strings.Title(strings.ToLower(strings.TrimSpace(name))),
		Symbol: strings.ToUpper(strings.TrimSpace(symbol)),
	}

	// continues if at least one of the filters is not empty
	var validate bool = false
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		validate = true
	} else {
		z := "[cryptovote.validation] at least on filter shouldn't be empty"
		err = errors.New(z)
		return retrievedCryptoVote, err
	}

	if validate {
		filter, err := utils.MarshalFilterCryptoVoteToBsonFilter(filterCryptoVote)
		if err != nil {
			return retrievedCryptoVote, err
		}

		// uses function from package dao
		retrievedCryptoVote, err = bo.ImplDAO.FindOne(filter)
		if err != nil {
			z := "[cryptovote.mongodb] Problems using bo.cryptoVoteDAO.FindOne: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}
	}
	return retrievedCryptoVote, err
}

/*
	UpdateOneCryptoVoteByFilter uses a filter to find one item and set newData to it before update database.
	Return a model.CryptoVote updated or nil, if filter mismatch.

	filterCryptoVote = model.FilterCryptoVote{
		Name: "Bitcoin",
		Symbol: "",
	}

	newCryptoData := model.CryptoVote{
		Name:   "FormiCOIN",
		Symbol:  "FORMFORMFORMFORM",
		Qtd_Upvote: 0,
		Qtd_Downvote: 0,
	}
*/
func (bo CryptoVoteBO) UpdateOneCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote, cryptoNewData model.CryptoVote) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote

	cryptoNewData.Name = strings.Title(strings.ToLower(strings.TrimSpace(cryptoNewData.Name)))
	cryptoNewData.Symbol = strings.ToUpper(strings.TrimSpace(cryptoNewData.Symbol))
	cryptoNewData.Sum = cryptoNewData.Qtd_Upvote - cryptoNewData.Qtd_Downvote

	// uses function from package bo
	validate, err := bo.ValidateCryptoVote(cryptoNewData)

	if validate {
		filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
		filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

		retrievedCryptoVote, err = bo.retrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
		if err != nil {
			return retrievedCryptoVote, err
		}

		// continues if ID isn't zero, because newData are validated
		if !retrievedCryptoVote.Id.IsZero() {
			idFilter := bson.M{"_id": retrievedCryptoVote.Id}

			bsonCryptoNewData, err := utils.MarshalCryptoVoteToBsonFilter(cryptoNewData)
			if err != nil {
				return retrievedCryptoVote, err
			}

			newData := bson.M{
				"$set": bsonCryptoNewData,
			}

			// uses a filter to find one model.CryptoVote and set newData within
			retrievedCryptoVote, err = bo.ImplDAO.UpdateOne(idFilter, newData)
			if err != nil {
				z := "[cryptovote.mongodb] Problems using bo.cryptoVoteDAO.UpdateOne: " + err.Error()
				err = errors.New(z)
				return retrievedCryptoVote, err
			}
		}
	}
	return retrievedCryptoVote, err
}

/*
	DeleteAllCryptoVoteByFilter uses filter to find many itens and delete all of them.
	Return number of itens deleted. Zero indicates filter mismatch.

	filterCryptoVote = model.FilterCryptoVote{
		Name: "Bitcoin",
		Symbol: "",
	}
*/
func (bo CryptoVoteBO) DeleteAllCryptoVoteByFilter(filterCryptoVote model.FilterCryptoVote) (int64, error) {
	var filter = bson.M{}
	var deletedCount int64 = 0
	var err error

	// creates a filter using args
	filterCryptoVote.Name = strings.Title(strings.ToLower(strings.TrimSpace(filterCryptoVote.Name)))
	filterCryptoVote.Symbol = strings.ToUpper(strings.TrimSpace(filterCryptoVote.Symbol))

	// continues if at least one of the filters is not empty
	var validate bool = false
	if validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Name) || validateCryptoVoteArgumentNotEmpty(filterCryptoVote.Symbol) {
		validate = true
	} else {
		z := "[cryptovote.validation] at least on filter shouldn't be empty"
		err = errors.New(z)
		return deletedCount, err
	}

	if validate {
		filter, err = utils.MarshalFilterCryptoVoteToBsonFilter(filterCryptoVote)
		if err != nil {
			return deletedCount, err
		}

		// uses function from package dao
		deletedCount, err = bo.ImplDAO.DeleteMany(filter)
		if err != nil {
			z := "Problems using bo.cryptoVoteDAO.DeleteMany: " + err.Error()
			err = errors.New(z)
			return deletedCount, err
		}
	}
	return deletedCount, err
}

func (bo CryptoVoteBO) DeleteAllCryptoVote() (int64, error) {
	var filter = bson.M{}
	var deletedCount int64 = 0
	var err error

	// uses function from package dao
	deletedCount, err = bo.ImplDAO.DeleteMany(filter)
	if err != nil {
		z := "Problems using bo.cryptoVoteDAO.DeleteMany: " + err.Error()
		err = errors.New(z)
		return deletedCount, err
	}
	return deletedCount, err
}

/*
	AddUpVote perform an addition of a Upvote attribute from model.CryptoVote, using a filter.
	Return a model.CryptoVote stored at database or nil, if filter mismatch.
*/
func (bo CryptoVoteBO) AddUpVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote
	var err error

	// uses internal function
	retrievedCryptoVote, err = bo.retrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
	if err != nil {
		return retrievedCryptoVote, err
	}

	// perform math before send data do DAO object
	if !retrievedCryptoVote.Id.IsZero() {
		newQtd_Upvote := retrievedCryptoVote.Qtd_Upvote + 1
		newSum := newQtd_Upvote - retrievedCryptoVote.Qtd_Downvote
		newSumAbsolute := utils.Abs(newSum)
		typeVote := "qtd_upvote"

		retrievedCryptoVote, err = bo.updateVote(retrievedCryptoVote, typeVote, newQtd_Upvote, newSum, newSumAbsolute)
	}
	return retrievedCryptoVote, err
}

/*
	AddDownVote perform an addition of a DownVote attribute from model.CryptoVote, using a filter.
	Return a model.CryptoVote stored at database or nil, if filter mismatch.
*/
func (bo CryptoVoteBO) AddDownVote(filterCryptoVote model.FilterCryptoVote) (model.CryptoVote, error) {
	var retrievedCryptoVote model.CryptoVote
	var err error

	// uses internal function
	retrievedCryptoVote, err = bo.retrieveOneCryptoVote(filterCryptoVote.Name, filterCryptoVote.Symbol)
	if err != nil {
		return retrievedCryptoVote, err
	}

	// perform calculation before send data do DAO object
	if !retrievedCryptoVote.Id.IsZero() {
		newQtd_Downvote := retrievedCryptoVote.Qtd_Downvote + 1
		newSum := retrievedCryptoVote.Qtd_Upvote - newQtd_Downvote
		newSumAbsolute := utils.Abs(newSum)
		typeVote := "qtd_downvote"

		retrievedCryptoVote, err = bo.updateVote(retrievedCryptoVote, typeVote, newQtd_Downvote, newSum, newSumAbsolute)
		if err != nil {
			z := "[cryptovote.mongodb] Problems using CryptoVoteBO.updateVote: " + err.Error()
			err = errors.New(z)
			return retrievedCryptoVote, err
		}
	}
	return retrievedCryptoVote, err
}

func (bo CryptoVoteBO) updateVote(retrievedCryptoVote model.CryptoVote, typeVote string, newQtd int64, newSum int64, newSumAbsolute int64) (model.CryptoVote, error) {
	var err error

	// creates filter
	filter := bson.M{"_id": retrievedCryptoVote.Id}
	newData := bson.M{
		"$set": bson.M{
			typeVote:       newQtd,
			"sum":          newSum,
			"sum_absolute": newSumAbsolute,
		},
	}

	retrievedCryptoVote, err = bo.ImplDAO.UpdateOne(filter, newData)
	if err != nil {
		z := "[cryptovote.mongodb] Problems using bo.cryptoVoteDAO.UpdateOne: " + err.Error()
		err = errors.New(z)
		return retrievedCryptoVote, err
	}
	return retrievedCryptoVote, err
}

/*
	SumaryAllCryptoVote perform search sorted by Sum, using pageSize to limit slice lenght.
	Returns a slice of []model.SumaryVote stored at database or nil, if filter mismatch.
*/
func (bo CryptoVoteBO) SumaryAllCryptoVote(pageSize int64) ([]model.SumaryCryptoVote, error) {
	var sumaryCryptoVotes []model.SumaryCryptoVote
	var retrievedCryptoVotes []model.CryptoVote
	var err error

	const ZERO int64 = 0
	const DEFAULT_PAGE_SIZE int64 = 10
	if pageSize == ZERO {
		// uses default value = 10
		pageSize = DEFAULT_PAGE_SIZE
	}

	filter, err := utils.MarshalFilterCryptoVoteToBsonFilter(utils.LoadOneNewEmptyFilterCryptoVote())
	if err != nil {
		return sumaryCryptoVotes, err
	}

	// orderType ascending = 1 / descending = -1
	orderType := -1
	sortFieldName := "sum_absolute"
	retrievedCryptoVotes, err = bo.ImplDAO.FindManyLimitedSortedByField(filter, pageSize, sortFieldName, orderType)
	if err != nil {
		z := "[cryptovote.mongodb] Problems at dao.FindManyCryptoVoteLimitedSortedByField: " + err.Error()
		err = errors.New(z)
		return sumaryCryptoVotes, err
	}

	t := int64(len(retrievedCryptoVotes))
	if t > pageSize {
		z := fmt.Sprintf("[cryptovote.mongodb] Total itens (%d) > then page size (%d)", t, pageSize)
		err = errors.New(z)
		return nil, err
	}

	sumaryCryptoVotes = nil
	for _, cryptoVote := range retrievedCryptoVotes {
		var sumary model.SumaryCryptoVote
		sumary.Token = cryptoVote.Name + "/" + cryptoVote.Symbol
		sumary.SumAbsolute = cryptoVote.SumAbsolute

		var sumType string
		if cryptoVote.Qtd_Upvote == cryptoVote.Qtd_Downvote {
			sumType = "Equal"
		} else {
			if cryptoVote.Qtd_Upvote > cryptoVote.Qtd_Downvote {
				sumType = "Up vote"
			} else {
				sumType = "Down vote"
			}
		}
		sumary.SumType = sumType
		sumaryCryptoVotes = append(sumaryCryptoVotes, sumary)
	}
	return sumaryCryptoVotes, err
}
