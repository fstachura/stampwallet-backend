package api

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	api "github.com/StampWallet/backend/internal/api/models"
	"github.com/StampWallet/backend/internal/database"
	acc "github.com/StampWallet/backend/internal/database/accessors"
	. "github.com/StampWallet/backend/internal/database/accessors/mocks"
	"github.com/StampWallet/backend/internal/managers"
	. "github.com/StampWallet/backend/internal/managers/mocks"
	. "github.com/StampWallet/backend/internal/testutils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lithammer/shortuuid/v4"
	"github.com/stretchr/testify/require"
)

func GetBusinessHandlers(ctrl *gomock.Controller) *BusinessHandlers {
	commonMockUserAuthorizedAccessor := NewMockUserAuthorizedAccessor(ctrl)
	commonMockBusinessAuthorizedAccessor := NewMockBusinessAuthorizedAccessor(ctrl)

	return &BusinessHandlers{
		businessManager:    NewMockBusinessManager(ctrl),
		transactionManager: NewMockTransactionManager(ctrl),
		itemDefinitionHandlers: ItemDefinitionHandlers{
			itemDefinitionManager:      NewMockItemDefinitionManager(ctrl),
			userAuthorizedAcessor:      commonMockUserAuthorizedAccessor,
			businessAuthorizedAccessor: commonMockBusinessAuthorizedAccessor,
			logger:                     log.Default(),
		},
		userAuthorizedAcessor:         commonMockUserAuthorizedAccessor,
		businessAuthorizedAccessor:    commonMockBusinessAuthorizedAccessor,
		authorizedTransactionAccessor: NewMockAuthorizedTransactionAccessor(ctrl),
		logger:                        log.Default(),
	}
}

func setupBusinessHandlersPostAccount() (
	w *httptest.ResponseRecorder,
	context *gin.Context,
	testBusinessUser *database.User,
	testBusiness *database.Business,
	testBusinessDetails *managers.BusinessDetails,
	respBodyExpected *api.PostBusinessAccountResponse,
) {
	testBusinessUser = GetDefaultUser()
	testBusiness = GetDefaultBusiness(testBusinessUser)
	testBusinessDetails = &managers.BusinessDetails{
		Name:        testBusiness.Name,
		Description: testBusiness.Description,
		Address:     testBusiness.Address,
		// GPSCoordinates: testBusiness.GPSCoordinates, TODO GPSCoordinates to string
		NIP:       testBusiness.NIP,
		KRS:       testBusiness.KRS,
		REGON:     testBusiness.REGON,
		OwnerName: testBusiness.OwnerName,
	}

	payload := api.PostBusinessAccountRequest{
		Name:    testBusiness.Name,
		Address: testBusiness.Address,
		// GpsCoordinates: testBusiness.GPSCoordinates, TODO GPSCoordinates to string
		Nip:       testBusiness.NIP,
		Krs:       testBusiness.KRS,
		Regon:     testBusiness.REGON,
		OwnerName: testBusiness.OwnerName,
	}
	payloadJson, _ := json.Marshal(payload)

	gin.SetMode(gin.TestMode)
	w = httptest.NewRecorder()

	context = NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/account").
		SetUser(testBusinessUser).
		SetMethod("POST").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected = &api.PostBusinessAccountResponse{
		PublicId:      testBusiness.PublicId,
		BannerImageId: testBusiness.BannerImageId,
		IconImageId:   testBusiness.IconImageId,
	}

	return w, context, testBusinessUser, testBusiness, testBusinessDetails, respBodyExpected
}

func TestBusinessHandlersPostAccountOk(t *testing.T) {
	w, context, testBusinessUser, testBusiness, testBusinessDetails, respBodyExpected := setupBusinessHandlersPostAccount()

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.businessManager.(*MockBusinessManager).
		EXPECT().
		Create(
			gomock.Eq(testBusinessUser),
			gomock.Eq(testBusinessDetails),
		).
		Return(
			testBusiness,
			nil,
		)

	handler.postAccount(context)

	respBody, respCode, respParseErr := ExtractResponse[api.PostBusinessAccountResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(200), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPostAccountNok_ChkBiz_InvalidNip(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	defaultBusiness := GetDefaultBusiness(testBusinessUser)
	defaultBusiness.NIP = "some invalid NIP"
	testBusinessDetails := &managers.BusinessDetails{
		Name:        defaultBusiness.Name,
		Description: defaultBusiness.Description,
		Address:     defaultBusiness.Address,
		// GPSCoordinates: defaultBusiness.GPSCoordinates, TODO GPSCoordinates to string
		NIP:       defaultBusiness.NIP,
		KRS:       defaultBusiness.KRS,
		REGON:     defaultBusiness.REGON,
		OwnerName: defaultBusiness.OwnerName,
	}

	payload := api.PostBusinessAccountRequest{
		Name:    testBusinessDetails.Name,
		Address: testBusinessDetails.Address,
		// GpsCoordinates: testBusinessDetails.GPSCoordinates, TODO GPSCoordinates to string
		Nip:       testBusinessDetails.NIP,
		Krs:       testBusinessDetails.KRS,
		Regon:     testBusinessDetails.REGON,
		OwnerName: testBusinessDetails.OwnerName,
	}
	payloadJson, _ := json.Marshal(payload)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/account").
		SetUser(testBusinessUser).
		SetMethod("POST").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected := &api.DefaultResponse{Status: api.INVALID_REQUEST}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.postAccount(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(400), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPostAccountNok_ChkBiz_InvalidKrs(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	defaultBusiness := GetDefaultBusiness(testBusinessUser)
	defaultBusiness.KRS = "some invalid KRS"
	testBusinessDetails := &managers.BusinessDetails{
		Name:        defaultBusiness.Name,
		Description: defaultBusiness.Description,
		Address:     defaultBusiness.Address,
		// GPSCoordinates: defaultBusiness.GPSCoordinates, TODO GPSCoordinates to string
		NIP:       defaultBusiness.NIP,
		KRS:       defaultBusiness.KRS,
		REGON:     defaultBusiness.REGON,
		OwnerName: defaultBusiness.OwnerName,
	}

	payload := api.PostBusinessAccountRequest{
		Name:    testBusinessDetails.Name,
		Address: testBusinessDetails.Address,
		// GpsCoordinates: testBusinessDetails.GPSCoordinates, TODO GPSCoordinates to string
		Nip:       testBusinessDetails.NIP,
		Krs:       testBusinessDetails.KRS,
		Regon:     testBusinessDetails.REGON,
		OwnerName: testBusinessDetails.OwnerName,
	}
	payloadJson, _ := json.Marshal(payload)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/account").
		SetUser(testBusinessUser).
		SetMethod("POST").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected := &api.DefaultResponse{Status: api.INVALID_REQUEST}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.postAccount(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(400), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPostAccountNok_ChkBiz_InvalidRegon(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	defaultBusiness := GetDefaultBusiness(testBusinessUser)
	defaultBusiness.REGON = "some invalid REGON"
	testBusinessDetails := &managers.BusinessDetails{
		Name:        defaultBusiness.Name,
		Description: defaultBusiness.Description,
		Address:     defaultBusiness.Address,
		// GPSCoordinates: defaultBusiness.GPSCoordinates, TODO GPSCoordinates to string
		NIP:       defaultBusiness.NIP,
		KRS:       defaultBusiness.KRS,
		REGON:     defaultBusiness.REGON,
		OwnerName: defaultBusiness.OwnerName,
	}

	payload := api.PostBusinessAccountRequest{
		Name:    testBusinessDetails.Name,
		Address: testBusinessDetails.Address,
		// GpsCoordinates: testBusinessDetails.GPSCoordinates, TODO GPSCoordinates to string
		Nip:       testBusinessDetails.NIP,
		Krs:       testBusinessDetails.KRS,
		Regon:     testBusinessDetails.REGON,
		OwnerName: testBusinessDetails.OwnerName,
	}
	payloadJson, _ := json.Marshal(payload)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/account").
		SetUser(testBusinessUser).
		SetMethod("POST").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected := &api.DefaultResponse{Status: api.INVALID_REQUEST}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.postAccount(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(400), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPostAccountNok_UniqBiz(t *testing.T) {
	w, context, testBusinessUser, _, testBusinessDetails, _ := setupBusinessHandlersPostAccount()

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.businessManager.(*MockBusinessManager).
		EXPECT().
		Create(
			gomock.Eq(testBusinessUser),
			gomock.Eq(testBusinessDetails),
		).
		Return(
			nil,
			managers.BusinessAlreadyExists,
		)

	respBodyExpected := api.DefaultResponse{Status: api.CONFLICT}

	handler.postAccount(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(409), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func setupBusinessHandlersGetAccountInfo() (
	w *httptest.ResponseRecorder,
	context *gin.Context,
	testBusinessUser *database.User,
	testBusiness *database.Business,
	respBodyExpected *api.GetBusinessAccountResponse,
) {
	testBusinessUser = GetDefaultUser()
	testBusiness = GetDefaultBusiness(testBusinessUser)
	testBusiness.MenuItems = []database.MenuItem{
		{
			BusinessId: testBusiness.ID,
			FileId:     shortuuid.New(),
		},
	}
	testItemDef := GetDefaultItem(testBusiness)
	testBusiness.ItemDefinitions = []database.ItemDefinition{
		*testItemDef,
	}

	gin.SetMode(gin.TestMode)
	w = httptest.NewRecorder()

	context = NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/info").
		SetUser(testBusinessUser).
		SetMethod("GET").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		Context

	respBodyExpected = &api.GetBusinessAccountResponse{
		PublicId: testBusiness.PublicId,
		Name:     testBusiness.Name,
		Address:  testBusiness.Address,
		// GpsCoordinates: testBusiness.GPSCoordinates, TODO GPSCoordinates
		BannerImageId: testBusiness.BannerImageId,
		IconImageId:   testBusiness.IconImageId,
		MenuImageIds: []string{
			testBusiness.MenuItems[0].FileId,
		},
		ItemDefinitions: []api.ItemDefinitionApiModel{
			{
				PublicId:    testItemDef.PublicId,
				Name:        testItemDef.Name,
				Price:       int32(testItemDef.Price),
				Description: testItemDef.Description,
				ImageId:     testItemDef.ImageId,
				StartDate:   &testItemDef.StartDate.Time,
				EndDate:     &testItemDef.EndDate.Time, // TODO: Should be ptr?
				MaxAmount:   int32(testItemDef.MaxAmount),
				Available:   testItemDef.Available,
			},
		},
	}

	return w, context, testBusinessUser, testBusiness, respBodyExpected
}

func TestBusinessHandlersGetAccountInfoOk(t *testing.T) {
	w, context, testBusinessUser, testBusiness, respBodyExpected := setupBusinessHandlersGetAccountInfo()

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.userAuthorizedAcessor.(*MockUserAuthorizedAccessor).
		EXPECT().
		Get(
			gomock.Eq(testBusinessUser),
			gomock.Eq(&database.Business{PublicId: testBusiness.PublicId}),
		).
		Return(
			testBusiness,
			nil,
		)

	handler.getAccountInfo(context)

	respBody, respCode, respParseErr := ExtractResponse[api.GetBusinessAccountResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(200), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersGetAccountInfoNok_NoBiz(t *testing.T) {
	w, context, testBusinessUser, _, _ := setupBusinessHandlersGetAccountInfo()

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.userAuthorizedAcessor.(*MockUserAuthorizedAccessor).
		EXPECT().
		Get(
			gomock.Eq(testBusinessUser),
			gomock.Eq(&database.Business{PublicId: "example of an invalid id"}),
		).
		Return(
			nil,
			acc.NotFound,
		)

	respBodyExpected := api.DefaultResponse{Status: api.NOT_FOUND}

	handler.getAccountInfo(context)

	respBody, respCode, respParseErr := ExtractResponse[api.GetBusinessAccountResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(404), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPatchAccountInfoOk(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	testBusiness := GetDefaultBusiness(testBusinessUser)

	payload := api.PatchBusinessAccountRequest{
		Name: "new business name",
	}
	payloadJson, _ := json.Marshal(payload)

	newBusinessDetails := &managers.BusinessDetails{
		Name: payload.Name,
	}

	// copying from ptr
	testBusinessVal := *testBusiness
	newBusiness := &testBusinessVal
	newBusiness.Name = newBusinessDetails.Name

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/info").
		SetUser(testBusinessUser).
		SetMethod("PATCH").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected := &api.DefaultResponse{Status: api.OK}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.userAuthorizedAcessor.(*MockUserAuthorizedAccessor).
		EXPECT().
		Get(
			gomock.Eq(testBusinessUser),
			gomock.Eq(&database.Business{PublicId: testBusiness.PublicId}),
		).
		Return(
			testBusiness,
			nil,
		)

	handler.businessManager.(*MockBusinessManager).
		EXPECT().
		ChangeDetails(
			gomock.Eq(testBusiness),
			gomock.Eq(newBusinessDetails),
		).
		Return(
			newBusiness,
			nil,
		)

	handler.patchAccountInfo(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(200), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

func TestBusinessHandlersPatchAccountInfoNok_InvBiz(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	testBusiness := GetDefaultBusiness(testBusinessUser)

	payload := api.PatchBusinessAccountRequest{
		Name: "new business name",
	}
	payloadJson, _ := json.Marshal(payload)

	newBusinessDetails := &managers.BusinessDetails{
		Name: payload.Name,
	}

	// copying from ptr
	testBusinessVal := *testBusiness
	newBusiness := &testBusinessVal
	newBusiness.Name = newBusinessDetails.Name

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/info").
		SetUser(testBusinessUser).
		SetMethod("PATCH").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	respBodyExpected := &api.DefaultResponse{Status: api.NOT_FOUND}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.userAuthorizedAcessor.(*MockUserAuthorizedAccessor).
		EXPECT().
		Get(
			gomock.Eq(testBusinessUser),
			gomock.Eq(&database.Business{PublicId: testBusiness.PublicId}),
		).
		Return(
			nil,
			acc.NotFound,
		)

	handler.patchAccountInfo(context)

	respBody, respCode, respParseErr := ExtractResponse[api.DefaultResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(404), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

// FUTURE: TestBusinessHandlersPatchAccountInfoNok_ChkBiz for code coverage,
// For now TestBusinessHandlersGetAccountInfoNok_ChkBiz is enough.

func TestBusinessHandlersGetTransactionOk(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	testTransactionCode := "0123456789"
	testBusiness := GetDefaultBusiness(testBusinessUser)
	testVcard := GetTestVirtualCard(nil, testBusinessUser, testBusiness)
	testItemDef := GetDefaultItem(testBusiness)
	testOwnedItem := GetDefaultOwnedItem(testItemDef, testVcard)
	testTransaction, testTransactionDetails := GetTestTransaction(
		nil,
		testVcard,
		[]database.OwnedItem{*testOwnedItem},
	)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/transaction/"+testTransactionCode).
		SetUser(testBusinessUser).
		SetMethod("GET").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		Context

	respBodyExpected := &api.GetBusinessTransactionResponse{
		PublicId:      testTransaction.PublicId,
		VirtualCardId: int32(testVcard.ID),
		State:         testTransaction.State, // TODO: type mismatch?
		Items: []api.TransactionItemDetailApiModel{
			{
				PublicId:         testOwnedItem.PublicId,
				ItemDefinitionId: testItemDef.PublicId,
			},
		},
	}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	// TODO Mock translation itemId to OwnedItem - accessor or manager?

	handler.authorizedTransactionAccessor.(*MockAuthorizedTransactionAccessor).
		EXPECT().
		GetForBusiness(
			gomock.Eq(testBusinessUser),
			gomock.Eq(testTransactionCode),
		).
		Return(
			testTransaction,
			testTransactionDetails,
			nil,
		)

	handler.getTransaction(context)

	respBody, respCode, respParseErr := ExtractResponse[api.GetBusinessTransactionResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(200), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq
}

// FUTURE: Rainy day scenarios for GetTransaction endpoint

func TestBusinessHandlersPostTransactionOk(t *testing.T) {
	testBusinessUser := GetDefaultUser()
	testBusiness := GetDefaultBusiness(testBusinessUser)
	testVcard := GetTestVirtualCard(nil, testBusinessUser, testBusiness)
	testItemDef := GetDefaultItem(testBusiness)
	testOwnedItem := GetDefaultOwnedItem(testItemDef, testVcard)
	testTransaction, testTransactionDetails := GetTestTransaction(
		nil,
		testVcard,
		[]database.OwnedItem{*testOwnedItem},
	)

	payload := api.PostBusinessTransactionRequest{
		AddedPoints: 10,
		ItemActions: []api.ItemActionApiModel{
			{
				ItemId: testOwnedItem.PublicId,
				Action: api.REDEEMED,
			},
		},
	}
	payloadJson, _ := json.Marshal(payload)

	testItemsWithAction := []managers.ItemWithAction{
		{
			Item:   testOwnedItem,
			Action: database.ActionTypeEnum(payload.ItemActions[0].Action), // TODO: Good conversion?
		},
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	context := NewTestContextBuilder(w).
		SetDefaultUrl().
		SetEndpoint("/business/transaction/"+testTransaction.Code).
		SetUser(testBusinessUser).
		SetMethod("POST").
		SetHeader("Content-Type", "application/json").
		SetDefaultToken().
		SetBody(payloadJson).
		Context

	transactionFinalized := &database.Transaction{
		PublicId:      testTransaction.PublicId,
		VirtualCardId: testTransaction.VirtualCardId,
		Code:          testTransaction.Code,
		State:         database.TransactionStateFinished,
		AddedPoints:   uint(payload.AddedPoints),
	}

	// test env prep
	ctrl := gomock.NewController(t)
	handler := GetBusinessHandlers(ctrl)

	handler.authorizedTransactionAccessor.(*MockAuthorizedTransactionAccessor).
		EXPECT().
		GetForBusiness(
			gomock.Eq(testBusinessUser),
			gomock.Eq(testTransaction.Code),
		).
		Return(
			testTransaction,
			testTransactionDetails,
			nil,
		)

		// TODO Mock translation itemId to OwnedItem - accessor or manager?

	handler.transactionManager.(*MockTransactionManager).
		EXPECT().
		Finalize(
			gomock.Eq(testTransaction),
			gomock.Eq(testItemsWithAction),
			gomock.Eq(payload.AddedPoints),
		).
		Return(
			transactionFinalized,
			nil,
		)

	handler.getTransaction(context)

	respBodyExpected := api.DefaultResponse{Status: api.OK}
	respBody, respCode, respParseErr := ExtractResponse[api.GetBusinessTransactionResponse](w)

	require.Nilf(t, respParseErr, "Failed to parse JSON response")
	require.Equalf(t, respCode, int(200), "Response returned unexpected status code")
	require.Truef(t, MatchEntities(respBodyExpected, respBody), "Response returned unexpected body contents")
	// TODO: test MatchEntities and gomock.Eq

}