using media_house_api;
using Microsoft.AspNetCore.Http.Features;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
string allowOriginsPolicy = "allowOrigins";

builder.Services.AddCors(options =>
{
    options.AddPolicy(name: allowOriginsPolicy,
    policy =>
        {
            policy.WithOrigins(
                AppSettings.LOCAL_HOST_PRODUCTION_DOMAIN,
                AppSettings.LOCAL_HOST_TESTING_DOMAIN
            );
            policy.AllowAnyHeader();
            policy.AllowAnyMethod();
        });
});

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseRouting();

app.UseCors(allowOriginsPolicy);

app.UseAuthorization();

app.MapControllers();

app.Run();
